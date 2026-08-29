package contour

import (
	"fmt"
	"os"
	"path"

	"github.com/wandb/server/pkg/dependency"
	"github.com/wandb/server/pkg/images"
	"github.com/wandb/server/pkg/kubernetes"
)

func manifestURL(version string) string {
	return fmt.Sprintf(
		"https://raw.githubusercontent.com/projectcontour/contour/release-%s/examples/render/contour.yaml",
		version,
	)
}

func NewPackage(version string, dest string) dependency.Package {
	return &ContourPackage{version, dest}
}

// ContourPackage is a package for installing contour. Contour is a Kubernetes
// ingress controller for Lyft's Envoy proxy.
type ContourPackage struct {
	version string
	dest    string
}

func (p ContourPackage) path() string {
	pa := path.Join(p.dest, p.Name(), p.version)
	os.MkdirAll(pa, 0755)
	return pa
}

func (p ContourPackage) Download() error {
	err := dependency.HTTPDownloadAndSave(
		manifestURL(p.version),
		path.Join(p.path(), "kube-contour.yml"),
	)
	if err != nil {
		return err
	}

	f, _ := os.ReadFile(path.Join(p.path(), "kube-contour.yml"))
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

func (p ContourPackage) Install() error {
	panic("unimplemented")
}

func (p ContourPackage) Name() string {
	return "contour"
}

func (p ContourPackage) Version() string {
	return p.version
}
