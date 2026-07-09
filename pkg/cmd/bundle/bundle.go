package bundle

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"os"
	"path/filepath"
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
	packages = append(
		packages,
		config.KubernetesImages()...,
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

func CreateBundle(dest string) error {
	srcs := []string{
		config.Config.Dir,
		"installer",
		"configs",
		"config.yaml",
		"install.sh",
		"LICENSE",
		"SECURITY.md",
	}

	destFile, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer destFile.Close()

	gzipWriter := gzip.NewWriter(destFile)
	defer gzipWriter.Close()

	tarWriter := tar.NewWriter(gzipWriter)
	defer tarWriter.Close()

	for _, src := range srcs {
		err := filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			header, err := tar.FileInfoHeader(info, info.Name())
			if err != nil {
				return err
			}
			header.Name = filepath.ToSlash(path)

			if err := tarWriter.WriteHeader(header); err != nil {
				return err
			}

			// Write file content
			if !info.Mode().IsRegular() {
				return nil
			}

			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()

			_, err = io.Copy(tarWriter, file)
			return err
		})

		if err != nil {
			return err
		}
	}

	return nil
}
