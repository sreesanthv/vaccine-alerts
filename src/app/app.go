package app

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/buger/jsonparser"
	"github.com/sreesanthv/vaccine-alerts/src/notification"
)

type App struct {
	Notifiers []notification.Notifier
	AppConf   *AppConf
}

type AppConf struct {
	CowinUrl       string
	CowinDistricts string
	AlertDays      int
	FirstDoseOnly  bool
	SecondDoseOnly bool
}

func (a *AppConf) GetDistrictIDs() []int {
	d := []int{}
	s := strings.Split(a.CowinDistricts, ",")
	for _, val := range s {
		dId, err := strconv.Atoi(val)
		if err != nil {
			continue
		}
		d = append(d, dId)
	}
	return d
}

func NewApp(conf *AppConf, not []notification.Notifier) *App {
	app := &App{
		Notifiers: not,
		AppConf:   conf,
	}
	return app
}

func (a *App) Start() {
	log.Println("Started fetching free slots")
	base := time.Now()

	count := 0
	for _, districtId := range a.AppConf.GetDistrictIDs() {
		for i := 0; i < a.AppConf.AlertDays; i++ {
			dt := base.AddDate(0, 0, i)
			date := dt.Format("02-01-2006")

			url := fmt.Sprintf("%s?district_id=%d&date=%s", a.AppConf.CowinUrl, districtId, date)
			res, err := http.Get(url)
			resBuf := new(bytes.Buffer)
			io.Copy(resBuf, res.Body)

			if err != nil {
				log.Println("Error connecting Cowin API:", err)
				continue
			} else if res.StatusCode != 200 {
				log.Println("Cowin API reponse not ok:", resBuf.String())
				continue
			}

			jsonparser.ArrayEach(resBuf.Bytes(), func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
				name, _ := jsonparser.GetString(value, "name")
				district, _ := jsonparser.GetString(value, "district_name")
				date, _ := jsonparser.GetString(value, "date")
				vaccine, _ := jsonparser.GetString(value, "vaccine")
				minAgeLimit, _ := jsonparser.GetInt(value, "min_age_limit")
				capacityDose1, _ := jsonparser.GetInt(value, "available_capacity_dose1")
				capacityDose2, _ := jsonparser.GetInt(value, "available_capacity_dose2")
				fee, _ := jsonparser.GetString(value, "fee")
				blockName, _ := jsonparser.GetString(value, "block_name")

				if capacityDose1 == 0 && capacityDose2 == 0 {
					return
				} else if a.AppConf.FirstDoseOnly && capacityDose1 == 0 {
					return
				} else if a.AppConf.SecondDoseOnly && capacityDose2 == 0 {
					return
				}

				feeFlt, _ := strconv.ParseFloat(fee, 32)
				if feeFlt == 0 {
					fee = "*Free*"
				} else {
					fee = "Fee:\t" + fee
				}

				content := []string{
					strings.Join([]string{"*" + date, name + "*"}, "\t"),
					strings.Join([]string{"\t", vaccine, fee}, "\t"),
					strings.Join([]string{"\t", fmt.Sprintf("Age: %d", minAgeLimit)}, "\t"),
					strings.Join([]string{"\t", fmt.Sprintf("D1: %d\tD2: %d", capacityDose1, capacityDose2)}, "\t"),
					strings.Join([]string{"\t", blockName, district}, "\t"),
				}

				for _, notifier := range a.Notifiers {
					notifier.Notify(content)
				}
				count++
			}, "sessions")

		}
	}
	log.Println("Fetching free slots completed. Count:", count)

	if count == 0 {
		for _, notifier := range a.Notifiers {
			notifier.Notify([]string{"No free slots available"})
		}
	}
}
