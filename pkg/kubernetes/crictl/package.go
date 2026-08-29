package crictl

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"runtime"

	"github.com/wandb/server/pkg/dependency"
)

const GithubRepo = "https://github.com/kubernetes-sigs/cri-tools"

func DownloadURL(version string) string {
	return fmt.Sprintf(
		"%s/releases/download/v%s/crictl-v%s-linux-%s.tar.gz",
		GithubRepo,
		version,
		version,
		runtime.GOARCH,
	)
}

func NewPackage(version string, dest string) dependency.Package {
	return &CrictlPackage{version, dest}
}

type CrictlPackage struct {
	version string
	dest string
}

func (p CrictlPackage) path() string {
	pa := path.Join(p.dest, p.Name(), p.version)
	os.MkdirAll(pa, 0755)
	return pa
}

func (p CrictlPackage) Version() string {
	return p.version
}

func (p CrictlPackage) Install() error {
	tar := path.Join(p.path(), "crictl.tar.gz")
	return exec.Command("tar", "-xzf", tar, "-C", "/usr/local/bin").Run()
}

func (p CrictlPackage) Download() error {
	return dependency.HTTPDownloadAndSave(
		DownloadURL(p.version),
		path.Join(p.path(), "crictl.tar.gz"),
	)
}

func (p CrictlPackage) Name() string {
	return "circtl"
}