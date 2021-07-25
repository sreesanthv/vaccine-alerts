package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/sreesanthv/vaccine-alerts/src/app"
	"github.com/sreesanthv/vaccine-alerts/src/notification"
)

var rootCmd = &cobra.Command{
	Use:   "vaccine-alerts",
	Short: "Fetch alerts from Cowin and send alerts",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func init() {
	viper.AutomaticEnv()
}

func validateMandatoryEnv(envs []string) {
	for _, env := range envs {
		if val := viper.GetString(env); val == "" {
			log.Fatalf("Please set value for Env variable, %s", env)
		}
	}
}

func StartBatch(notifiers []notification.Notifier) {
	validateMandatoryEnv([]string{"COWIN_URL", "COWIN_DISTRICT_IDS", "ALERT_DAYS"})
	app := app.NewApp(&app.AppConf{
		CowinUrl:        viper.GetString("COWIN_URL"),
		CowinDistricts:  viper.GetString("COWIN_DISTRICT_IDS"),
		AlertDays:       viper.GetInt("ALERT_DAYS"),
		FirstDoseOnly:   viper.GetBool("COWIN_FIRST_DOSE_ONLY"),
		SecondDoseOnly:  viper.GetBool("COWIN_SECOND_DOSE_ONLY"),
		FreeVaccineOnly: viper.GetBool("COWIN_FREE_VACCINE_ONLY"),
	}, notifiers)

	app.Start()
}
