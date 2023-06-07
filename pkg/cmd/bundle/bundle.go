package bundle

import (
	"sync"

	"github.com/pterm/pterm"
	"github.com/wandb/server/pkg/config"
	"github.com/wandb/server/pkg/dependency"
)

func DownloadAllPackages() {
	packages := append(
		config.KubernetesPackages(),
		config.KubernetesAddonPackages()...,
	)

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
				pterm.Error.Printf("Failed to download %s (v%s): %e\n", p.Name(), p.Version(), err)
			}
			progressbar.Increment()
			wg.Done()
			pterm.Error.PrintOnError(err)
		}(pkg)
	}
	wg.Wait()

	progressbar.Stop()
}

func DownloadImages() {

}
