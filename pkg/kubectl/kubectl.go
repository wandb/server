package kubectl

import (
	"os"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func GetClientset() (*rest.Config, *kubernetes.Clientset, error) {
	kubeconfigPath := os.Getenv("KUBECONFIG")
    if kubeconfigPath == "" {
        kubeconfigPath = os.Getenv("HOME") + "/.kube/config"
    }
    config, err := clientcmd.BuildConfigFromFlags("", kubeconfigPath)
    if err != nil {
        return nil, nil, err
    }

	cs, err := kubernetes.NewForConfig(config)
	return config, cs, err
}

func IsConnectedToCluster() bool {
	_, _, err := GetClientset()
	return err == nil
}

