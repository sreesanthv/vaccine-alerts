package cmd

import (
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
	validateMandatoryEnv([]string{"SLACK_WEBHOOK_URL"})

	notifier := notification.NewSlackNotifier(viper.GetString("SLACK_WEBHOOK_URL"))

	app := app.NewApp(&app.AppConf{
		CowinUrl:       viper.GetString("COWIN_URL"),
		CowinDistricts: viper.GetString("COWIN_DISTRICT_IDS"),
		AlertDays:      viper.GetInt("ALERT_DAYS"),
		FirstDoseOnly:  viper.GetBool("COWIN_FIRST_DOSE_ONLY"),
		SecondDoseOnly: viper.GetBool("COWIN_SECOND_DOSE_ONLY"),
	}, []notification.Notifier{notifier})
	app.Start()
}
