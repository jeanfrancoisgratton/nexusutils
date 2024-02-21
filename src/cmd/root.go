// nxrmutils
// src/cmd/root.go

package cmd

import (
	"github.com/spf13/cobra"
	"nxrmutils/env"
	"nxrmutils/helpers"
	"os"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "nxrmutils",
	Short:   "Various tools to manage users, blobs, repos, etc",
	Version: "1.00.00-0 (2024.02.19)",
	Long: `This tools allows you to manage components in Nexus Repository Manager.
You can manage blob stores, repositories, users, etc`,
}

var clCmd = &cobra.Command{
	Use:     "changelog",
	Aliases: []string{"cl"},
	Short:   "Shows changelog",
	Run: func(cmd *cobra.Command, args []string) {
		helpers.ChangeLog()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.DisableAutoGenTag = true
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.AddCommand(clCmd)

	rootCmd.PersistentFlags().StringVarP(&env.EnvConfigFile, "env", "e", "defaultEnv.json", "Default environment configuration file; this is a per-user setting.")
	//upCmd.Flags().Int8VarP(&exec.IndexNumber, "index", "i", 0, "Index of repository; this is zero-based.")
}
