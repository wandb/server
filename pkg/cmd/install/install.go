package install

import (
	"sync"

	"github.com/pterm/pterm"
	"github.com/wandb/server/pkg/config"
	"github.com/wandb/server/pkg/dependency"
)

func InstallKubernetes() {
	packages := config.KubernetesPackages()

	progressbar, _ := pterm.DefaultProgressbar.
		WithTotal(len(packages)).
		WithTitle("Installing kubernetes").
		Start()

	wg := sync.WaitGroup{}
	wg.Add(len(packages))
	for _, pkg := range packages {
		go func(install dependency.Package) {
			install.Install()
			wg.Done()
		}(pkg)
	}
	wg.Wait()
	progressbar.Stop()
}

func InstallKubernetesAddons() {
	packages := config.KubernetesAddonPackages()

	progressbar, _ := pterm.DefaultProgressbar.
		WithTotal(len(packages)).
		WithTitle("Installing kubernetes").
		Start()

	wg := sync.WaitGroup{}
	wg.Add(len(packages))
	for _, pkg := range packages {
		go func(install dependency.Package) {
			install.Install()
			wg.Done()
		}(pkg)
	}
	wg.Wait()
	progressbar.Stop()
}

func InstallWandbOperator() {
}
