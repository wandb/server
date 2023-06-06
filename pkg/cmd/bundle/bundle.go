package bundle

import (
	"os"
	"sync"

	"github.com/pterm/pterm"
	"github.com/wandb/server/pkg/dependency"
	"github.com/wandb/server/pkg/kubernetes/addons/flannel"
	"github.com/wandb/server/pkg/kubernetes/helm"
)

func DownloadPackages() {
	downloadDir := "./packages"
	os.MkdirAll(downloadDir, 0755)
	packages := []dependency.Package{
		// containerd.NewPackage("1.7.2", downloadDir),
		// crictl.NewPackage("1.27.0", downloadDir),
		// runc.NewPackage("1.1.7", downloadDir),
		// cni.NewPackage("1.3.0", downloadDir),
		// conntrack.NewPackage("1.4.6-2", downloadDir),
		// kubeadm.NewPackage("1.27.2", downloadDir),
		// kubectl.NewPackage("1.27.2", downloadDir),
		// kubelet.NewPackage("1.27.2", downloadDir),
		flannel.NewPackage("0.22.0", downloadDir),
		helm.NewPackage("3.12.0", downloadDir),
	}

	progressbar, _ := pterm.DefaultProgressbar.
		WithTotal(len(packages)).
		WithTitle("Downloading packages").
		Start()
	
	wg := sync.WaitGroup{}
	wg.Add(len(packages))
	for _, pkg := range packages {
		go func(p dependency.Package) {
			pterm.Info.Printf("Downloading %s (v%s)\n", p.Name(), p.Version())
			err := p.Download()
			if err == nil {
				pterm.Success.Printf("Downloaded %s (v%s)\n", p.Name(), p.Version())
			} else {
				pterm.Error.Printf("Failed to download %s (v%s): %w\n", p.Name(), p.Version(), err)
			}
			progressbar.Increment()
			wg.Done()
			pterm.Error.PrintOnError(err)
		}(pkg)
	}
	wg.Wait()

	progressbar.Stop()
}
