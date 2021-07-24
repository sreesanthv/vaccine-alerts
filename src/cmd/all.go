package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/sreesanthv/vaccine-alerts/src/app"
	"github.com/sreesanthv/vaccine-alerts/src/notification"
)

var cmdAllAlert = &cobra.Command{
	Use:   "all",
	Short: "Send alerts through all mediums",
	Run:   sendAllAlerts,
}

func init() {
	rootCmd.AddCommand(cmdAllAlert)
}

func sendAllAlerts(cmd *cobra.Command, args []string) {
	validateMandatoryEnv([]string{"SLACK_WEBHOOK_URL", "TELEGRAM_BOT_TOKEN", "TELEGRAM_CHAT_ID"})
	notifiers := []notification.Notifier{
		notification.NewSlackNotifier(viper.GetString("SLACK_WEBHOOK_URL")),
		notification.NewTelegramNotifier(viper.GetString("TELEGRAM_BOT_TOKEN"), viper.GetString("TELEGRAM_CHAT_ID")),
	}
	app := app.NewApp(&app.AppConf{
		CowinUrl:       viper.GetString("COWIN_URL"),
		CowinDistricts: viper.GetString("COWIN_DISTRICT_IDS"),
		AlertDays:      viper.GetInt("ALERT_DAYS"),
		FirstDoseOnly:  viper.GetBool("COWIN_FIRST_DOSE_ONLY"),
		SecondDoseOnly: viper.GetBool("COWIN_SECOND_DOSE_ONLY"),
	}, notifiers)
	app.Start()
}
