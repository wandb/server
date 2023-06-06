package kubectl

import (
	"fmt"
	"os"
	"path"
	"runtime"

	"github.com/wandb/server/pkg/dependency"
	"github.com/wandb/server/pkg/files"
)

func DownloadURL(version string) string {
	return fmt.Sprintf(
		"https://dl.k8s.io/release/v%s/bin/linux/%s/kubectl",
		version,
		runtime.GOARCH,
	)
}

func NewPackage(version string, dest string) dependency.Package {
	return &KubectlPackage{version, dest}
}

type KubectlPackage struct {
	version string
	dest string
}

func (p KubectlPackage) path() string {
	pa := path.Join(p.dest, p.Name(), p.version)
	os.MkdirAll(pa, 0755)
	return pa
}


func (p KubectlPackage) Version() string {
	return p.version
}

func (p KubectlPackage) Install() error {
	binary := path.Join(p.path(), "kubectl")
	err := files.CopyFile(binary, "/usr/local/bin/kubectl")
	if err != nil {
		return err
	}

	return os.Chmod("/usr/local/bin/kubectl", 0755)
}

func (p KubectlPackage) Download() error {
	return dependency.HTTPDownloadAndSave(
		DownloadURL(p.version),
		path.Join(p.path(), "kubectl"),
	)
}

func (p KubectlPackage) Name() string {
	return "kubectl"
}