// Copyright © 2020 The pf9ctl authors

package cmd

import (
	"go.uber.org/zap"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a resource",
	Long: `Use the create command to create cluster, context, support bundle and
	other resources`,
	Run: func(cmd *cobra.Command, args []string) {
		zap.S().Info("Create called")
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}
