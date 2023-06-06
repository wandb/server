package kubernetes

import (
	"fmt"
	"strings"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer/yaml"
)

func GetImagesFromManifest(data string) ([]string, error) {
	dec := yaml.NewDecodingSerializer(unstructured.UnstructuredJSONScheme)
	images := make([]string, 0)
	chunks := strings.Split(data, "---")

	for _, chunk := range chunks {
		chunk = strings.TrimSpace(chunk)
		if len(chunk) == 0 {
			continue
		}

		obj := &unstructured.Unstructured{}
		_, _, err := dec.Decode([]byte(chunk), nil, obj)
		if err != nil {
			return nil, fmt.Errorf("error decoding manifests: %v", err)
		}

		containers, _, _ := unstructured.NestedSlice(obj.Object, "spec", "template", "spec", "containers")
		for _, cnt := range containers {
			container := &corev1.Container{}
			err := runtime.DefaultUnstructuredConverter.FromUnstructured(cnt.(map[string]interface{}), container)
			if err != nil{
				return nil, fmt.Errorf("failed to convert unstructured container: %v", err)
			}
			images = append(images, container.Image)
		}

		initContainers, _, _ := unstructured.NestedSlice(obj.Object, "spec", "template", "spec", "initContainers")
		for _, cnt := range initContainers {
			initContainer := &corev1.Container{}
			err := runtime.DefaultUnstructuredConverter.FromUnstructured(cnt.(map[string]interface{}), initContainer)
			if err != nil {
				return nil, fmt.Errorf("failed to convert unstructured init container: %v", err)
			}
			images = append(images, initContainer.Image)
		}
	}

	return images, nil
}