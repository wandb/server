package crd

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/wandb/server/pkg/helm/values"
	"github.com/wandb/server/pkg/kubectl"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

const (
	Kind         = "WeightsAndBiases"
	Plural       = "weightsandbiases"
	Group        = "apps.wandb.com"
	GroupVersion = "v1"

	Name      = "wandb"
	Namespace = "default"
)

type WeightsAndBiases struct {
	APIVersion string `json:"apiVersion"`
	Kind       string `json:"kind"`
	Metadata   struct {
		Name      string            `json:"name"`
		Namespace string            `json:"namespace"`
		Labels    map[string]string `json:"labels"`
	} `json:"metadata"`
	Spec Spec `json:"spec"`
}

type Spec struct {
	Chart struct {
		Path string `json:"path"`
	} `json:"chart"`
	Values map[string]interface{} `json:"values"`
}

func NewWeightsAndBiases(chartPath string, vals values.Values) *WeightsAndBiases {
	return &WeightsAndBiases{
		APIVersion: fmt.Sprintf("%s/%s", Group, GroupVersion),
		Kind:       Kind,
		Metadata: struct {
			Name      string            `json:"name"`
			Namespace string            `json:"namespace"`
			Labels    map[string]string `json:"labels"`
		}{
			Name:      Name,
			Namespace: Namespace,
			Labels: map[string]string{
				"app.kubernetes.io/name":     Plural,
				"app.kubernetes.io/instance": Name,
			},
		},
		Spec: Spec{
			Chart: struct {
				Path string `json:"path"`
			}{
				Path: chartPath,
			},
			Values: vals,
		},
	}
}

func ApplyWeightsAndBiases(wb *WeightsAndBiases) error {
	_, cs, err := kubectl.GetClientset()
	if err != nil {
		return err
	}

	wandbJson, err := json.Marshal(wb)
	if err != nil {
		return err
	}

	gvr := schema.GroupVersionResource{
		Group:    Group,
		Version:  GroupVersion,
		Resource: Plural,
	}

	var unstructured unstructured.Unstructured
	if err := json.Unmarshal(wandbJson, &unstructured); err != nil {
		return err
	}
	_, err = cs.Resource(gvr).Namespace("default").Apply(context.Background(), "wandb", &unstructured, v1.ApplyOptions{
		FieldManager: "wandb-controller-manager",
		Force:        true,
	})

	return err
}
