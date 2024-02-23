package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/wandb/server/pkg/deployer"
	"github.com/wandb/server/pkg/helm"
	"github.com/wandb/server/pkg/images"
	"github.com/wandb/server/pkg/term/pkgm"
	"github.com/wandb/server/pkg/utils"
	"gopkg.in/yaml.v2"
)

func RootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "wbs",
	}

	return cmd
}

func downloadChartImages(
	url string,
	name string,
	version string,
	vals map[string]interface{},
) ([]string, error) {
	chartsDir := "bundle/charts"
	if err := os.MkdirAll(chartsDir, 0755); err != nil {
		return nil, err
	}

	chart, err := helm.DownloadChart(
		url,
		name,
		version,
		chartsDir,
	)
	if err != nil {
		return nil, err
	}

	runs, err := helm.GetRuntimeObjects(chart, vals)
	if err != nil {
		return nil, err
	}
	return helm.ExtractImages(runs), nil
}

func DownloadCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "download",
		Run: func(cmd *cobra.Command, args []string) {
			// TODO: make this download latest chart and use the latest
			// controller docker image by default the chart would download
			// latest and we should probably explicitly set the value
			fmt.Println("Downloading operator helm chart")
			operatorImgs, _ := downloadChartImages(
				"https://charts.wandb.ai",
				"operator",
				"1.1.0",
				map[string]interface{}{
					"image": map[string]interface{}{
						"tag": "1.10.1",
					},
				},
			)

			spec, err := deployer.GetChannelSpec("")
			if err != nil {
				panic(err)
			}

			yamlData, err := yaml.Marshal(spec)
			if err != nil {
				panic(err)
			}
			if err = os.WriteFile("bundle/spec.yaml", yamlData, 0644); err != nil {
				panic(err)
			}

			fmt.Println("Downloading wandb helm chart")
			// operator-wandb helm chart
			wandbImgs, _ := downloadChartImages(
				spec.Chart.URL,
				spec.Chart.Name,
				spec.Chart.Version,
				spec.Values,
			)

			imgs := utils.RemoveDuplicates(append(wandbImgs, operatorImgs...))
			if len(imgs) == 0 {
				fmt.Println("No images to download.")
				os.Exit(1)
			}

			cb := func(pkg string) {
				path := "bundle/images/" + strings.ReplaceAll(pkg, ":", "/")
				os.MkdirAll(path, 0755)
				err := images.Download(pkg, path+"/image.tgz")
				if err != nil {
					fmt.Println(err)
				}
			}
			fmt.Println("Deploying images")
			if _, err := pkgm.New(imgs, cb).Run(); err != nil {
				fmt.Println("Error running program:", err)
				os.Exit(1)
			}
		},
	}

	return cmd
}

func DeployCmd() *cobra.Command {
	var deployWithOperator bool
	var bundlePath string
	cmd := &cobra.Command{
		Use: "deploy",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}

	cmd.Flags().BoolVarP(&deployWithOperator, "operator", "o", true, "Deploy the system using the operator pattern.")
	cmd.Flags().StringVarP(&bundlePath, "bundle", "b", "", "Path to the bundle to deploy with.")

	cmd.Flags().MarkHidden("operator")

	return cmd
}

func main() {
	ctx := context.Background()
	cmd := RootCmd()
	cmd.AddCommand(DownloadCmd())
	cmd.AddCommand(DeployCmd())
	cmd.ExecuteContext(ctx)
}
