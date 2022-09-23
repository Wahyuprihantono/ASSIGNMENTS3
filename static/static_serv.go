package static

import (
	"ASSIGNMENTS3/models"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"text/template"
	"time"
)

func AutoReload() {
	for {
		water := RandomNumberWater()
		wind := RandomNumberWind()

		num := models.StatusAirAngin{}
		num.Status.Air = water
		num.Status.Angin = wind

		jsonData, err := json.Marshal(num)

		if err != nil {
			log.Fatal(err.Error())
		}
		err = ioutil.WriteFile("data.json", jsonData, 0644)

		if err != nil {
			log.Fatal(err.Error())
		}
		time.Sleep(15 * time.Second)
	}
}

func ReloadWeb(w http.ResponseWriter, r *http.Request) {
	fileData, err := ioutil.ReadFile("data.json")

	if err != nil {
		log.Fatal(err.Error())
	}

	var statusData models.StatusAirAngin

	err = json.Unmarshal(fileData, &statusData)
	if err != nil {
		log.Fatal(err.Error())
	}

	waterVal := statusData.Status.Air
	windVal := statusData.Status.Angin

	var (
		waterStatus string
		windStatus  string
	)

	waterStatus = KondisiWater(waterVal)
	windStatus = KondisiWind(windVal)

	data := map[string]interface{}{
		"waterStatus": waterStatus,
		"windStatus":  windStatus,
		"waterHeight": waterVal,
		"windSpeed":   windVal,
	}

	tpl, err := template.ParseFiles("./asset/template.html")

	if err != nil {
		log.Fatal(err.Error())
	}

	tpl.Execute(w, data)

}
