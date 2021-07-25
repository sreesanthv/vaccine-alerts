package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
	StartBatch([]notification.Notifier{notifier})
}
