package conntrack

import (
	"fmt"
	"os"
	"path"
	"runtime"

	"github.com/wandb/server/pkg/dependency"
	"github.com/wandb/server/pkg/linux/debian"
)

func DownloadURL(version string) string {
	return fmt.Sprintf(
		"https://deb.debian.org/debian/pool/main/c/conntrack-tools/conntrack_%s_%s.deb",
		version,
		runtime.GOARCH,
	)
}

func NewPackage(version string, dest string) dependency.Package {
	return &ConntrackPackage{version, dest}
}

type ConntrackPackage struct {
	version string
	dest string
}

func (p ConntrackPackage) path() string {
	pa := path.Join(p.dest, p.Name(), p.version)
	os.MkdirAll(pa, 0755)
	return pa
}

func (p *ConntrackPackage) Version() string {
	return p.version
}

func (p *ConntrackPackage) Install() error {
	return debian.InstallPackage(p.dest)
}

func (p *ConntrackPackage) Download() error {
	return dependency.HTTPDownloadAndSave(
		DownloadURL(p.version),
		path.Join(p.path(), "conntrack.deb"),
	)
}

func (p *ConntrackPackage) Name() string {
	return "conntrack"
}
