package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path"

	"github.com/spf13/cobra"
	"github.com/wandb/server/cmds/wbs/deploy"
	"github.com/wandb/server/pkg/deployer"
	"github.com/wandb/server/pkg/helm"
	"github.com/wandb/server/pkg/helm/values"
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
			fmt.Println("Downloading operator helm chart")
			operatorImgs, _ := downloadChartImages(
				helm.WandbHelmRepoURL,
				helm.WandbOperatorChart,
				"", // empty version means latest
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
				path := "bundle/images/" + pkg
				os.MkdirAll(path, 0755)
				err := images.Download(pkg, path+"/image.tgz")
				if err != nil {
					fmt.Println(err)
				}
			}

			if _, err := pkgm.New(imgs, cb).Run(); err != nil {
				fmt.Println("Error deploying:", err)
				os.Exit(1)
			}
		},
	}

	return cmd
}

func DeployCmd() *cobra.Command {
	var deployWithHelm bool
	var bundlePath string
	var namespace string
	releaseName := "wandb"
	var valuesPath string
	var chartPath string

	cmd := &cobra.Command{
		Use: "deploy",
		Run: func(cmd *cobra.Command, args []string) {
			homedir, err := os.UserHomeDir()
			if err != nil {
				fmt.Printf("could not find home dir: %v", err)
				os.Exit(1)
			}

			chartsDir := path.Join(homedir, ".wandb", "charts")
			os.MkdirAll(chartsDir, 0755)

			if deployWithHelm {
				spec := deploy.GetChannelSpec()
				// Merge user values with spec values
				vals := spec.Values
				if localVals, err := values.FromYAMLFile(valuesPath); err == nil {
					if finalVals, err := vals.Merge(localVals); err != nil {
						vals = finalVals
					}
				}

				if chartPath == "" {
					fmt.Println("Downloading W&B chart from", spec.Chart.URL)
					chartPath = deploy.DownloadHelmChart(
						spec.Chart.URL, spec.Chart.Name, spec.Chart.Version, chartsDir)
				}
				chart := deploy.LoadChart(chartPath)
				if _, err := json.Marshal(vals.AsMap()); err != nil {
					panic(err)
				}
				deploy.DeployChart(namespace, releaseName, chart, vals.AsMap())
				os.Exit(0)
			}
		},
	}

	cmd.Flags().BoolVarP(&deployWithHelm, "helm", "", false, "Deploy the system using the helm (not recommended).")
	cmd.Flags().StringVarP(&bundlePath, "bundle", "b", "", "Path to the bundle to deploy with.")
	cmd.Flags().StringVarP(&valuesPath, "values", "v", "", "Values file to apply to the helm chart yaml.")
	cmd.Flags().StringVarP(&namespace, "namespace", "n", "wandb", "Namespace to deploy into.")
	cmd.Flags().StringVarP(&namespace, "chart", "c", "", "Path to W&B helm chart.")

	cmd.Flags().MarkHidden("helm")

	return cmd
}

func main() {
	ctx := context.Background()
	cmd := RootCmd()
	cmd.AddCommand(DownloadCmd())
	cmd.AddCommand(DeployCmd())
	cmd.ExecuteContext(ctx)
}
