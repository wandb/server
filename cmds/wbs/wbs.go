package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"path"

	"github.com/spf13/cobra"
	"github.com/wandb/server/cmds/wbs/deploy"
	"github.com/wandb/server/pkg/deployer"
	"github.com/wandb/server/pkg/helm"
	"github.com/wandb/server/pkg/helm/values"
	"github.com/wandb/server/pkg/kubectl"
	"github.com/wandb/server/pkg/utils"
	"gopkg.in/yaml.v2"
	"k8s.io/apimachinery/pkg/runtime/schema"
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

			yamlData, err := yaml.Marshal(spec)
			if err != nil {
				panic(err)
			}
			if err = os.WriteFile("bundle/spec.yaml", yamlData, 0644); err != nil {
				panic(err)
			}

			// cb := func(pkg string) {
			// 	path := "bundle/images/" + pkg
			// 	os.MkdirAll(path, 0755)
			// 	err := images.Download(pkg, path+"/image.tgz")
			// 	if err != nil {
			// 		fmt.Println(err)
			// 	}
			// }

			// if _, err := pkgm.New(imgs, cb).Run(); err != nil {
			// 	fmt.Println("Error deploying:", err)
			// 	os.Exit(1)
			// }
		},
	}

	return cmd
}

type Charts map[string]struct {
	Binary string
}

func base64EncodeFile(filePath string) (string, error) {

	fileContents, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(fileContents), nil
}

func deployOperator(chartsDir string, namespace string, releaseName string) error {
	operatorChartPath := deploy.DownloadHelmChart(
		helm.WandbHelmRepoURL, helm.WandbOperatorChart, "", chartsDir,
	)

	operatorChart := deploy.LoadChart(operatorChartPath)
	if operatorChart == nil {
		return fmt.Errorf("Could not load chart")
	}

	wandbChartPath := deploy.DownloadHelmChart(
		helm.WandbHelmRepoURL, helm.WandbChart, "", chartsDir,
	)

	wandbChartBinary, err := base64EncodeFile(wandbChartPath)
	if err != nil {
		return err
	}

	operatorValues := values.Values{}
	operatorValues.SetValue("airgapped", true)

	operatorValues.SetValue("charts", Charts{
		helm.WandbChart: {
			Binary: wandbChartBinary,
		},
	})

	deploy.DeployChart(namespace, releaseName, operatorChart, operatorValues.AsMap())

	return nil
}

type LocalSpec struct {
	Chart struct {
		Path string `json:"path" yaml:"path"`
	} `json:"chart" yaml:"chart"`
	Values values.Values `json:"values" yaml:"values"`
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

			spec := deploy.GetChannelSpec()
			// Merge user values with spec values
			vals := spec.Values
			if localVals, err := values.FromYAMLFile(valuesPath); err == nil {
				if finalVals, err := vals.Merge(localVals); err != nil {
					vals = finalVals
				}
			}

			if deployWithHelm {
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

			if err := deployOperator(chartsDir, namespace, releaseName); err != nil {
				fmt.Println("Error deploying operator:", err)
				os.Exit(1)
			}

			// deploy weightsandbiases crd
			_, cs, err := kubectl.GetClientset()
			if err != nil {
				fmt.Println("Error getting clientset:", err)
				os.Exit(1)
			}

			gvr := schema.GroupVersionResource{
				Group:    "apps.wandb.com",
				Version:  "v1",
				Resource: "weightsandbiases",
			}

			localSpec := LocalSpec{
				Chart: struct {
					Path string `json:"path" yaml:"path"`
				}{
					Path: fmt.Sprintf("charts/%s", helm.WandbChart),
				},
				Values: vals,
			}
			// cs.AppsV1().RESTClient().Post().Resource("deployments").Namespace("default").Body(`{"apiVersion":"apps/v1","kind":"Deployment","metadata":{"name":"nginx"},"spec":{"replicas":3,"selector":{"matchLabels":{"app":"nginx"}},"template":{"metadata":{"labels":{"app":"nginx"}},"spec":{"containers":[{"name":"nginx","image":"nginx:1.7.9","ports":[{"containerPort":80}]}]}}}}`).Do(context.Background())

			// if bundlePath == "" {
			// 	bundlePath = "bundle/"
			// }

			// wandbChartPattern := fmt.Sprintf("charts/%s-*.tgz", helm.WandbChart)

			// wandbFileMatches, err := filepath.Glob(bundlePath + wandbChartPattern)
			// if err != nil {
			// 	fmt.Println("Error finding wandb chart:", err)
			// 	os.Exit(1)
			// }

			// if len(wandbFileMatches) == 0 {
			// 	fmt.Println("Wandb chart not found in bundle")
			// 	os.Exit(1)
			// }

			// wandbChartPath := wandbFileMatches[0]
			// wandbChartPathTrimmed := strings.TrimPrefix(wandbChartPath, bundlePath)

			// fmt.Println(wandbChartPathTrimmed)

			// // deploy operator
			// operatorValues := values.Values{}
			// operatorValues.SetValue("airgapped", true)

			// // charts := []ZippedChart{
			// // 	{
			// // 		Path:  wandbChartPathTrimmed,
			// // 		Chart: helm.WandbChart,
			// // 	},
			// // }

			// operatorChartPattern := fmt.Sprintf("charts/%s-*.tgz", helm.WandbOperatorChart)
			// operatorFileMatches, err := filepath.Glob(bundlePath + operatorChartPattern)
			// if err != nil {
			// 	fmt.Println("Error finding operator chart:", err)
			// 	os.Exit(1)
			// }

			// if len(operatorFileMatches) == 0 {
			// 	fmt.Println("Operator chart not found in bundle")
			// 	os.Exit(1)
			// }

			// var operatorChartPath string

			// for _, file := range operatorFileMatches {
			// 	if strings.Contains(file, "operator-wandb") {
			// 		continue
			// 	}

			// 	operatorChartPath = file
			// }

			// operatorValues.SetValue("charts", charts)

			// fmt.Println("Deploying operator chart from", operatorChartPath)
			// fmt.Println("Charts to mount", wandbChartPath)

			// operatorChart := deploy.LoadChart(operatorChartPath)
			// if operatorChart == nil {
			// 	fmt.Println("Could not load chart")
			// 	os.Exit(1)
			// }

			// deploy.DeployChart("wandb", "operator", operatorChart, operatorValues.AsMap())

			// deploy weightsandbiases crd

			os.Exit(0)
		},
	}

	cmd.Flags().BoolVarP(&deployWithHelm, "helm", "", false, "Deploy the system using the helm (not recommended).")
	cmd.Flags().StringVarP(&bundlePath, "bundle", "b", "", "Path to the bundle to deploy with.")
	cmd.Flags().StringVarP(&valuesPath, "values", "v", "", "Values file to apply to the helm chart yaml.")
	cmd.Flags().StringVarP(&namespace, "namespace", "n", "wandb", "Namespace to deploy into.")
	cmd.Flags().StringVarP(&chartPath, "chart", "c", "", "Path to W&B helm chart.")

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
