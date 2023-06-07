package images

import (
	"fmt"
	"os/exec"

	"github.com/wandb/server/pkg/kubernetes/containerd"
	"github.com/wandb/server/pkg/kubernetes/docker"
)

// Download downloads an image and saves it to a file.
func Download(image string, filename string) error {
	_, err := exec.LookPath("docker")
	if err == nil {
		return docker.DownloadImage(image, filename)
	}

	if containerd.IsInstalled() {
		return containerd.DownloadImage(image, filename)
	}

	return fmt.Errorf("no supported container runtime found")
}
