package task

import (
	redis "github.com/go-redis/redis/v9"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "task",
	Short: "Simple cli task manager",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}
var (
	cfgFile string
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConf)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "conf.yml", "config file (default $DIR/conf.yaml)")

	// due to all commands need access to redis we connect to it in root command
	rdb := redis.NewClient(&redis.Options{
		Addr: viper.GetString("db_url"),
	})
	rootCmd.AddCommand(
		NewAddCmd(rdb),
		NewListCommand(rdb),
		NewDoCommand(rdb),
	)
}

func initConf() {
	viper.SetConfigFile(cfgFile)

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}
