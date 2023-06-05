package images

import (
	"fmt"
	"os/exec"

	"github.com/wandb/server/pkg/kubernetes/containerd"
	"github.com/wandb/server/pkg/kubernetes/docker"
)

// ImageDownloadAndSave downloads an image and saves it to a file.
func DownloadImageAndSave(image string, tag string, filename string) error {
	_, err := exec.LookPath("docker")
	if err != nil {
		return docker.DownloadImageWithDocker(image, tag, filename)
	}

	if containerd.IsInstalled() {
		return containerd.DownloadImage(image, tag, filename)
	}

	return fmt.Errorf("no supported container runtime found")
}
