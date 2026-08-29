package runc

import (
	"fmt"
	"os"
	"path"
	"runtime"

	"github.com/wandb/server/pkg/dependency"
	"github.com/wandb/server/pkg/files"
)

const GithubRepo = "https://github.com/opencontainers/runc"

func DownloadURL(version string) string {
	return fmt.Sprintf(
		"%s/releases/download/v%s/runc.%s",
		GithubRepo,
		version,
		runtime.GOARCH,
	)
}

func NewPackage(version string, dest string) dependency.Package {
	return &RuncPackage{version, dest}
}

type RuncPackage struct {
	version string
	dest string
}

func (p RuncPackage) path() string {
	pa := path.Join(p.dest, p.Name(), p.version)
	os.MkdirAll(pa, 0755)
	return pa
}

func (p RuncPackage) Version() string {
	return p.version
}

func (p RuncPackage) Install() error {
	binary := path.Join(p.path(), "runc")
	err := files.CopyFile(binary, "/usr/local/sbin/runc")
	if err != nil {
		return err
	}

	return os.Chmod("/usr/local/sbin/runc", 0755)
}

func (p RuncPackage) Download() error {
	return dependency.HTTPDownloadAndSave(
		DownloadURL(p.version),
		path.Join(p.path(), "runc"),
	)
}

func (p RuncPackage) Name() string {
	return "runc"
}