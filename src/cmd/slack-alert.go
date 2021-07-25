package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
	StartBatch([]notification.Notifier{notifier})
}
