package main

import (
	"log"

	"github.com/clems4ever/anytype-backup-node/internal/backupnode"
	"github.com/spf13/cobra"
)

var configPathFlag string

var rootCmd = cobra.Command{
	Use:   "anytype-backup-node",
	Short: "Start an anytype backup node",
}

var generateNetconfCmd = cobra.Command{
	Use:   "generate-netconf",
	Short: "Generate the network configuration of an anytype backup node",
	Run: func(cmd *cobra.Command, args []string) {
		backupnode.GenerateConfig(configPathFlag)
	},
}

var bootstrapNodeCmd = cobra.Command{
	Use:   "bootstrap",
	Short: "Bootstrap a backup node",
	Run: func(cmd *cobra.Command, args []string) {
		backupnode.Bootstrap(cmd.Context(), configPathFlag)
	},
}

var initializeNodeCmd = cobra.Command{
	Use:   "initialize",
	Short: "Initialize the backup node",
	Run: func(cmd *cobra.Command, args []string) {
		backupnode.Initialize(cmd.Context(), configPathFlag)
	},
}

func main() {
	generateNetconfCmd.Flags().StringVarP(&configPathFlag, "config", "c", "", "path to the config file")
	generateNetconfCmd.MarkFlagRequired("c")

	rootCmd.AddCommand(&generateNetconfCmd)

	bootstrapNodeCmd.Flags().StringVarP(&configPathFlag, "config", "c", "", "path to the config file")
	bootstrapNodeCmd.MarkFlagRequired("c")
	rootCmd.AddCommand(&bootstrapNodeCmd)

	initializeNodeCmd.Flags().StringVarP(&configPathFlag, "config", "c", "", "path to the config file")
	initializeNodeCmd.MarkFlagRequired("c")
	rootCmd.AddCommand(&initializeNodeCmd)

	err := rootCmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
