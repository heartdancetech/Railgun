package cmd

import (
	"github.com/MisakaSystem/LastOrder/core"
	"github.com/MisakaSystem/LastOrder/logger"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// Used for flags.
	cfgFile     string
	userLicense string
	//databaseName string
	//host string
	//port int
	//user,password string
	//outdir string

	rootCmd = &cobra.Command{
		Use:   "LastOrder",
		Short: "",
		Long:  "",
		PreRun: func(cmd *cobra.Command, args []string) {
		},
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			logger.Init("debug")
			g := core.New()
			_ = g.Run()
		},
	}
)

// Execute executes the root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// TODO below params for test
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra.yaml)")
	rootCmd.PersistentFlags().StringVarP(&userLicense, "license", "l", "mit", "name of license for the project")
	rootCmd.PersistentFlags().Bool("viper", true, "use Viper for configuration")
	rootCmd.PersistentFlags().String("keo", "keo", "keo")
	_ = viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))
	_ = viper.BindPFlag("keo", rootCmd.PersistentFlags().Lookup("keo"))
	_ = viper.BindEnv("keo", "TEST_KEO")
	viper.SetDefault("author", "NAME HERE <EMAIL ADDRESS>")
}

func initConfig() {
	viper.SetEnvPrefix("test")
	viper.AutomaticEnv()
}
