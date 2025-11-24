package images

import (
	"context"
	"fmt"
	"os/exec"

	"github.com/containers/image/v5/copy"
	"github.com/containers/image/v5/directory"
	"github.com/containers/image/v5/docker"
	"github.com/containers/image/v5/signature"
	"github.com/containers/image/v5/types"
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

// Pull pulls an image from a registry and saves it to a directory not tarzipped.
func Pull(imageName string, dest string) error {
	ctx := context.Background()
	sysCtx := &types.SystemContext{
		OSChoice: "linux",
		ArchitectureChoice: "amd64",
	}

	srcRef, err := docker.ParseReference("//" + imageName)
	if err != nil {
		return fmt.Errorf("parsing source reference: %w", err)
	}

	destRef, err := directory.NewReference(dest)

	if err != nil {
		return fmt.Errorf("creating directory reference: %w", err)
	}

	policy := new(signature.Policy)
	policy.Default = []signature.PolicyRequirement{signature.NewPRInsecureAcceptAnything()}

	policyCtx, _ := signature.NewPolicyContext(policy)
	if err != nil {
		return fmt.Errorf("getting default policy context: %w", err)
	}
	defer policyCtx.Destroy()

	_, err = copy.Image(ctx, policyCtx, destRef, srcRef, &copy.Options{
		SourceCtx: sysCtx,
		DestinationCtx: sysCtx,
	})
	if err != nil {
		return fmt.Errorf("copying image: %w", err)
	}

	return nil
}
