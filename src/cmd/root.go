package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "vaccine-alerts",
	Short: "Fetch alerts from Cowin and send alerts",
}

var mandatoryEnvs = []string{"COWIN_URL", "COWIN_DISTRICT_IDS", "ALERT_DAYS"}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func init() {
	viper.AutomaticEnv()
	for _, env := range mandatoryEnvs {
		if val := viper.GetString(env); val == "" {
			log.Fatalf("Please set value for Env variable, %s", env)
		}
	}
}
