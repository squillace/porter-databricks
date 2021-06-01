package main

import (
	"github.com/spf13/cobra"
	"github.com/squillace/porter-databricks/pkg/databricks"
)

func buildUninstallCommand(m *databricks.Mixin) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "uninstall",
		Short: "Execute the uninstall functionality of this mixin",
		RunE: func(cmd *cobra.Command, args []string) error {
			return m.Execute()
		},
	}
	return cmd
}
