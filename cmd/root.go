package main

import (
	"context"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:               "octopus",
		DisableAutoGenTag: true,
		SilenceErrors:     true,
		SilenceUsage:      true,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.HelpFunc()(cmd, args)
		},
	}
)

func Execute(ctx context.Context) error {
	rootCmd.AddCommand(newServerCommand())
	return rootCmd.ExecuteContext(ctx)
}
