package helm

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"runtime"

	"github.com/wandb/server/pkg/dependency"
)

func DownloadURL(version string) string {
	return fmt.Sprintf(
		"https://get.helm.sh/helm-v%s-linux-%s.tar.gz",
		version,
		runtime.GOARCH,
	)
}

func NewPackage(version string, dist string) dependency.Package {
	return &HelmPackage{version, dist}
}

type HelmPackage struct {
	version string
	dest    string
}

func (p HelmPackage) path() string {
	pa := path.Join(p.dest, p.Name(), p.version)
	os.MkdirAll(pa, 0755)
	return pa
}

func (p HelmPackage) Download() error {
	return dependency.HTTPDownloadAndSave(
		DownloadURL(p.version),
		path.Join(p.path(), "helm.tar.gz"),
	)
}

func (p HelmPackage) Install() error {
	tar := path.Join(p.path(), "helm.tar.gz")
	return exec.Command("tar", "-xzf", tar, "-C", "/usr/local/bin").Run()
}

func (HelmPackage) Name() string {
	return "helm"
}

func (p HelmPackage) Version() string {
	return p.version
}
