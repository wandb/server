package flannel

import (
	"fmt"
	"os"
	"path"
	"strings"

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
	imgs, _ := kubernetes.GetImagesFromManifest(string(f))
	for _, image := range imgs {
		dir := path.Join(p.path(), strings.Split(image, ":")[0])
		os.MkdirAll(dir, 0755)
		images.Download(image, path.Join(dir, "image.tar"))
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
