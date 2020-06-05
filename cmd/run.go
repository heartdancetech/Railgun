package cmd

import (
	"github.com/MisakaSystem/LastOrder/core"
	"github.com/MisakaSystem/LastOrder/logger"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var runCmd = &cobra.Command{
	Use:     "run",
	Short:   "run",
	Long:    "run",
	Example: "run",
	Run: func(cmd *cobra.Command, args []string) {
		logger.Init("debug", viper.GetString("name"))
		g := core.New()
		_ = g.Run()
	},
}
