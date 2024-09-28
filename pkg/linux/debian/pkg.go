package debian

import (
	"fmt"
	"os/exec"
)

func InstallPackage(filepath string) error {
	cmd := exec.Command("dpkg", "-i", filepath)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to install .deb package: %v\nOutput: %s", err, string(output))
	}
	return nil
}