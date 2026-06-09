package images

import (
	"fmt"

	"github.com/wandb/server/pkg/kubernetes/containerd"
)

func LoadImage(filename string) error {
	if containerd.IsInstalled(){
		return containerd.LoadImage(filename)
	}
	return fmt.Errorf("no supported container runtime found")
}