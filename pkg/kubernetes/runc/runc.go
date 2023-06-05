package runc

import (
	"fmt"
	"os"
	"runtime"

	"github.com/pterm/pterm"
	"github.com/wandb/server/pkg/download"
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

func Download(version string, path string) error {
	pterm.Info.Printf("Downloading runc: v%s\n", version)
	return download.HTTPDownloadAndSave(DownloadURL(version), path)
}

func Install(file string) {
	pterm.Info.Printf("Installing runc from %s\n", file)
	files.CopyFile(file, "/usr/local/sbin/runc")
	os.Chmod("/usr/local/sbin/runc", 0755)
	pterm.Success.Println("Installed runc")
}