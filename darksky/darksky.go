package darksky

import (
	"fmt"
	"log"

	"../appconfig"

	"github.com/shawntoffel/darksky"
)

//Forecast retrieves weather forecast for given latitude and longitude
func Forecast(latitude float64, longitude float64) {
	client := darksky.New("e812d94640053e7d3f93bd055cfb246a")
	request := darksky.ForecastRequest{}

	request.Latitude = darksky.Measurement(latitude)
	request.Longitude = darksky.Measurement(longitude)
	request.Options = darksky.ForecastRequestOptions{Exclude: "hourly,minutely"}

	forecast, err := client.Forecast(request)
	if err != nil {
		log.Fatal(err)
	}

	println("Temperature for " + appconfig.Config.Home.City + ":")
	fmt.Printf("%.1f", forecast.Currently.Temperature)
	print("˚F\n")
}
