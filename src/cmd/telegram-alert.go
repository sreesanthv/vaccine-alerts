package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/sreesanthv/vaccine-alerts/src/app"
	"github.com/sreesanthv/vaccine-alerts/src/notification"
)

var cmdTelegramAlert = &cobra.Command{
	Use:   "telegram",
	Short: "Send alerts to telegram channel",
	Run:   sendTelegramAlerts,
}

func init() {
	rootCmd.AddCommand(cmdTelegramAlert)
}

func sendTelegramAlerts(cmd *cobra.Command, args []string) {
	validateMandatoryEnv([]string{"TELEGRAM_BOT_TOKEN", "TELEGRAM_CHAT_ID"})

	notifier := notification.NewTelegramNotifier(viper.GetString("TELEGRAM_BOT_TOKEN"), viper.GetString("TELEGRAM_CHAT_ID"))

	app := app.NewApp(&app.AppConf{
		CowinUrl:       viper.GetString("COWIN_URL"),
		CowinDistricts: viper.GetString("COWIN_DISTRICT_IDS"),
		AlertDays:      viper.GetInt("ALERT_DAYS"),
		FirstDoseOnly:  viper.GetBool("COWIN_FIRST_DOSE_ONLY"),
		SecondDoseOnly: viper.GetBool("COWIN_SECOND_DOSE_ONLY"),
	}, []notification.Notifier{notifier})
	app.Start()
}
