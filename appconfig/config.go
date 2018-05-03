package appconfig

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

//Home represents the user's default locale
type Home struct {
	City      string `json:"city"`
	State     string `json:"state"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

//Configuration represents the overall app configuration
type Configuration struct {
	Home Home `json:"home"`
}

//Config is the golbal configuration variable that holds all config values
var Config Configuration

//GetConfig reads config.json and stores the values in the config struct
func GetConfig() {
	configData, err := ioutil.ReadFile("./appconfig/config.json")
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(configData, &Config)
}

//SetHomeLocale sets the latitude and longitude of the user input
func SetHomeLocale(city string, state string, latitude string, longitude string) (string, error) {
	var message string
	Config.Home.City = city
	Config.Home.State = state
	Config.Home.Latitude = latitude
	Config.Home.Longitude = longitude

	c, err := json.Marshal(Config)
	if err != nil {
		message = "Failed to update home locale"
	}

	err = ioutil.WriteFile("./appconfig/config.json", c, 0644)
	if err != nil {
		message = "Failed to update home locale"
	}

	if message == "" {
		message = "Updated home locale successfully"
	}
	println(message)
	return message, nil
}
