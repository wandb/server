package images

import (
	"fmt"
	"os/exec"
)

func Download(image string, filename string) error {
	if _, err := exec.LookPath("docker"); err == nil {
		return DownloadUsingDocker(image, filename)
	}

	return fmt.Errorf("no supported container runtime found")
}

func DownloadUsingDocker(image string, filename string) error {
	cmdPull := exec.Command("docker", "pull", image)
	err := cmdPull.Run()
	if err != nil {
		return fmt.Errorf("failed to pull image using docker %s: %v", image, err)
	}

	cmdSave := exec.Command("docker", "save", image, "-o", filename)
	err = cmdSave.Run()
	if err != nil {
		return fmt.Errorf("failed to save image using docker %s as .tgz: %v", image, err)
	}

	return nil
}
