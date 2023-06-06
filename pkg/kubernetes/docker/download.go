package docker

import (
	"fmt"
	"os/exec"
)

func DownloadImage(image string, filename string) error {
	cmdPull := exec.Command("docker", "pull", image)
	err := cmdPull.Run()
	if err != nil {
		return fmt.Errorf("failed to pull docker image %s: %v", image, err)
	}

	cmdSave := exec.Command("docker", "save", image, "-o", filename)
	err = cmdSave.Run()
	if err != nil {
		return fmt.Errorf("failed to save docker image %s as .tgz: %v", image, err)
	}

	return nil
}

func IsInstalled() bool {
	_, err := exec.LookPath("docker")
	return err != nil
}