// nexusutils
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: src/cmd/blobCmd.go
// Original timestamp: 2024/02/21 17:17

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"nexusutils/blob"
)

var blobCmd = &cobra.Command{
	Use:     "blob",
	Aliases: []string{"environment"},
	Short:   "Blob store subcommands",
	Long: `Commands to handle blob stores.
You need to provide the subcommands: ls, create, rm.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("You need to provide one of the following subcommands: ls, create, or rm")
	},
}

var blobLsCmd = &cobra.Command{
	Use:     "ls",
	Aliases: []string{"list"},
	Short:   "List blob stores",
	Run: func(cmd *cobra.Command, args []string) {
		if err := blob.ListBlobStores(); err != nil {
			fmt.Println(err)
		}
	},
}
