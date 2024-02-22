package main

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/wandb/server/pkg/helm"
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
			helm.DownloadChart("https://charts.wandb.ai", "operator-wandb", "0.10.42")
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
