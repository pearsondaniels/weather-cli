package darksky

import (
	"fmt"
	"log"

	"github.com/shawntoffel/darksky"
)

//Forecast retrieves weather forecast for given latitude and longitude
func Forecast(city string, latitude float64, longitude float64) {
	client := darksky.New("e812d94640053e7d3f93bd055cfb246a")
	request := darksky.ForecastRequest{}

	request.Latitude = darksky.Measurement(latitude)
	request.Longitude = darksky.Measurement(longitude)
	request.Options = darksky.ForecastRequestOptions{Exclude: "hourly,minutely"}

	forecast, err := client.Forecast(request)
	if err != nil {
		log.Fatal(err)
	}

	println("Weather report for " + city + ":")

	print("Temperature: ")
	fmt.Printf("%.1f", forecast.Currently.Temperature)
	print("˚F\n")

	print("Precipitation: ")
	fmt.Printf("%.2f", forecast.Currently.PrecipIntensity*100)
	print("%\n")

	// print("Wind Speed: ")
	fmt.Printf("Wind Speed: %.1fmph\n", forecast.Currently.WindSpeed)
	// print("mph\n")

	// print("Humidity: ")
	fmt.Printf("Humidity: %.1f%%\n", forecast.Currently.Humidity*100)
	// print("%\n")
}
