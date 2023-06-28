package cmd

import "github.com/spf13/cobra"

func NewEnvCmd() *cobra.Command {
	envCmd := &cobra.Command{
		Use:   "env",
		Short: "A brief description of your application",
		Long:  `A longer description that spans multiple lines and likely contains`,
		Run: func(cmd *cobra.Command, args []string) {
			_ = cmd.Help() // Ignore error
		},
	}
	return envCmd
}
