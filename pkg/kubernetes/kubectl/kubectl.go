package kubectl

import (
	"fmt"
	"os"
	"runtime"

	"github.com/pterm/pterm"
	"github.com/wandb/server/pkg/download"
	"github.com/wandb/server/pkg/files"
)

func DownloadURL(version string) string {
	return fmt.Sprintf(
		"https://dl.k8s.io/release/v%s/bin/linux/%s/kubectl",
		version,
		runtime.GOARCH,
	)
}

func Download(version string, path string) error {
	pterm.Info.Printf("Downloading kubectl: v%s\n", version)
	return download.HTTPDownloadAndSave(DownloadURL(version), path)
}

func Install(binary string) {
	pterm.Info.Printf("Installing kubectl from %s\n", binary)
	err := files.CopyFile(binary, "/usr/local/bin/kubectl")
	pterm.Error.PrintOnError(err)
	err = os.Chmod("/usr/local/bin/kubectl", 0755)
	pterm.Error.PrintOnError(err)
}