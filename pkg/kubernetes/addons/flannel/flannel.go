package flannel

import (
	"fmt"
	"os"
	"path"

	"github.com/wandb/server/pkg/dependency"
	"github.com/wandb/server/pkg/images"
	"github.com/wandb/server/pkg/kubernetes"
)

const githubRepo = "https://github.com/flannel-io/flannel"

func manifestURL(version string) string {
	return fmt.Sprintf(
		"%s/releases/download/v%s/kube-flannel.yml",
		githubRepo,
		version,
	)
}

func NewPackage(version string, dest string) dependency.Package {
	return &FlannelPackage{version, dest}
}

// FlannelPackage is a package for installing flannel. Flannel is a Kubernetes
// network fabric for containers.
type FlannelPackage struct {
	version string
	dest    string
}

func (p FlannelPackage) path() string {
	pa := path.Join(p.dest, p.Name(), p.version)
	os.MkdirAll(pa, 0755)
	return pa
}

func (p FlannelPackage) Download() error {
	err := dependency.HTTPDownloadAndSave(
		manifestURL(p.version),
		path.Join(p.path(), "kube-flannel.yml"),
	)
	if err != nil {
		return err
	}

	f, _ := os.ReadFile(path.Join(p.path(), "kube-flannel.yml"))
	imgs, err := kubernetes.GetImagesFromManifest(string(f))
	if err != nil {
		return err
	}
	for _, image := range imgs {
		pkg := images.NewPackage(image, p.dest)
		if err = pkg.Download(); err != nil {
			return err
		}
	}
	return nil
}

func (p FlannelPackage) Install() error {
	panic("unimplemented")
}

func (p FlannelPackage) Name() string {
	return "flannel"
}

func (p FlannelPackage) Version() string {
	return p.version
}
