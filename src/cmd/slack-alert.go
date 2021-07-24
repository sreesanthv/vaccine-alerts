package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/sreesanthv/vaccine-alerts/src/app"
	"github.com/sreesanthv/vaccine-alerts/src/notification"
)

var cmdSlackAlert = &cobra.Command{
	Use:   "slack",
	Short: "Send alerts to slack channel",
	Run:   sendSlackAlerts,
}

func init() {
	rootCmd.AddCommand(cmdSlackAlert)
}

func sendSlackAlerts(cmd *cobra.Command, args []string) {
	slackUrl := viper.GetString("SLACK_WEBHOOK_URL")
	if slackUrl == "" {
		log.Fatal("Please set SLACK_WEBHOOK_URL env")
	}

	notifier := notification.NewSlackNotifier(slackUrl)

	app := app.NewApp(&app.AppConf{
		CowinUrl:       viper.GetString("COWIN_URL"),
		CowinDistricts: viper.GetString("COWIN_DISTRICT_IDS"),
		AlertDays:      viper.GetInt("ALERT_DAYS"),
	}, notifier)
	app.Start()
}
