package main

import (
	"context"
	"os"

	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"github.com/wandb/server/pkg/cmd/bundle"
	"github.com/wandb/server/pkg/cmd/install"
	"github.com/wandb/server/pkg/config"
	"github.com/wandb/server/pkg/kubernetes/kubeadm"
	"github.com/wandb/server/pkg/linux/swap"
)

func RootCmd() *cobra.Command {
	return &cobra.Command{
		Short: "W&B Server Installer CLI",
		Long:  `Weights & Biases tool for install and managing VM deployments`,
	}
}

func BundleCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "bundle",
		Short: "Creates an Airgap bunddle",
		Run: func(cmd *cobra.Command, args []string) {
			bundle.DownloadAllPackages()
		},
	}
}

func InstallCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "install",
		Short: "Runs the installer",
		Run: func(cmd *cobra.Command, args []string) {
			swap.MustSweepoff()

			bundle.DownloadAllPackages()
			bundle.DownloadImages()

			install.InstallKubernetes()

			kubeadm.Init()

			install.InstallKubernetesAddons()
			install.InstallWandbOperator()
		},
	}
}

func init() {
	pterm.EnableDebugMessages()
	pterm.EnableColor()
	pterm.DefaultInteractiveSelect.MaxHeight = 15
	pterm.Debug.Prefix.Text = "d"
	pterm.Debug.Prefix.Style = &pterm.ThemeDefault.DebugMessageStyle
	pterm.Info.Prefix.Text = "i"
	pterm.Info.Prefix.Style = &pterm.ThemeDefault.InfoMessageStyle
	pterm.Info.MessageStyle = &pterm.ThemeDefault.DefaultText
	pterm.Success.Prefix.Text = "✓"
	pterm.Success.Prefix.Style = &pterm.ThemeDefault.SuccessMessageStyle
	pterm.Warning.Prefix.Text = "!"
	pterm.Warning.Prefix.Style = &pterm.ThemeDefault.WarningMessageStyle
	pterm.Error.Prefix.Text = "✗"
	pterm.Error.Prefix.Style = &pterm.ThemeDefault.ErrorMessageStyle
	pterm.Fatal.Prefix.Text = "🤯"
	pterm.Fatal.Prefix.Style = &pterm.ThemeDefault.FatalMessageStyle
}

func main() {
	ctx := context.Background()
	cmd := RootCmd()

	os.MkdirAll(config.Config.Dir, 0755)

	cmd.AddCommand(InstallCommand())
	cmd.AddCommand(BundleCmd())

	cmd.ExecuteContext(ctx)
}
