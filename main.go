package main

import (
	"log"

	"github.com/clems4ever/anytype-backup-node/internal/backupnode"
	"github.com/spf13/cobra"
)

var configPathFlag string

var defaultConfigPath = "config.yml"

var rootCmd = cobra.Command{
	Use:   "anytype-backup-node",
	Short: "Start an anytype backup node",
}

var manualCmd = cobra.Command{
	Use:   "manual",
	Short: "Some commands that can be run manually if necessary.",
	Long:  "Some commands that can be run manually if you know what you are doing, otherwise let the bootstrap command handle it for you.",
}

var initCmd = cobra.Command{
	Use:   "init",
	Short: "Generate a default configuration file that can be edited before bootstraping the infrastructure",
	Run: func(cmd *cobra.Command, args []string) {
		backupnode.Init()
	},
}

var generateNetconfCmd = cobra.Command{
	Use:   "generate",
	Short: "Generate the configuration of an anytype backup node",
	Long: "Generate the configuration of an anytype backup node." +
		"This is only useful if you want to manually generate the configuration, " +
		"otherwise you can let the bootstrap command do it for you.",
	Run: func(cmd *cobra.Command, args []string) {
		backupnode.GenerateConfig(configPathFlag)
	},
}

var bootstrapNodeCmd = cobra.Command{
	Use:   "bootstrap",
	Short: "Bootstrap a backup node after you have generated a config file with the init command",
	Long:  "Bootstrap a backup node from the configuration file",
	Run: func(cmd *cobra.Command, args []string) {
		backupnode.Bootstrap(cmd.Context(), configPathFlag)
	},
}

var configureCmd = cobra.Command{
	Use:   "configure",
	Short: "Configure the infrastructure after it has been spawned",
	Long: "Configure the infrastructure after it has been spawned. " +
		"This is only useful if you have to do it manually, otherwise you can let " +
		"the bootstrap command handle it for you.",
	Run: func(cmd *cobra.Command, args []string) {
		backupnode.Initialize(cmd.Context(), configPathFlag)
	},
}

func main() {
	rootCmd.AddCommand(&initCmd)
	rootCmd.AddCommand(&manualCmd)

	generateNetconfCmd.Flags().StringVarP(&configPathFlag, "config", "c", defaultConfigPath, "path to the config file")
	manualCmd.AddCommand(&generateNetconfCmd)

	bootstrapNodeCmd.Flags().StringVarP(&configPathFlag, "config", "c", defaultConfigPath, "path to the config file")
	rootCmd.AddCommand(&bootstrapNodeCmd)

	configureCmd.Flags().StringVarP(&configPathFlag, "config", "c", defaultConfigPath, "path to the config file")
	manualCmd.AddCommand(&configureCmd)

	err := rootCmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
