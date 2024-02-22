package main

import (
	"context"
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/spf13/cobra"
	"github.com/wandb/server/pkg/deployer"
	"github.com/wandb/server/pkg/helm"
	"github.com/wandb/server/pkg/images"
)

func RootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "wbs",
	}

	return cmd
}

func DownloadCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "download",
		Run: func(cmd *cobra.Command, args []string) {
			spec, err := deployer.GetChannelSpec("")
			if err != nil {
				panic(err)
			}

			chart, _ := helm.DownloadChart(
				spec.Chart.URL, spec.Chart.Name, spec.Chart.Version)
			runs, _ := helm.GetRuntimeObjects(chart, spec.Values)
			imgs := helm.ExtractImages(runs)
			wg := sync.WaitGroup{}
			for _, image := range imgs {
				wg.Add(1)
				go func(image string) {
					fmt.Println("Downloading", image)
					path := "bundle/images/" + strings.ReplaceAll(image, ":", "/")
					os.MkdirAll(path, 0755)
					err := images.Download(image, path + "/image.tgz")
					if err != nil {
						fmt.Println(err)
					}
					wg.Done()
				}(image)
			}
			wg.Wait()
		},
	}

	return cmd
}

func DeployCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "deploy",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}

	return cmd
}

func main() {
	ctx := context.Background()
	cmd := RootCmd()
	cmd.AddCommand(DownloadCmd())
	cmd.AddCommand(DeployCmd())
	cmd.ExecuteContext(ctx)
}
