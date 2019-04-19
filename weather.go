package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/pearsondaniels/weather-cli/appconfig"
	"github.com/pearsondaniels/weather-cli/darksky"

	"github.com/manifoldco/promptui"
	cli "github.com/urfave/cli"
	"golang.org/x/net/context"
	"googlemaps.github.io/maps"
)

func main() {
	createApp()
}

func createApp() {
	app := cli.NewApp()
	appconfig.LoadConfig()

	app.Flags = []cli.Flag{
	// cli.StringFlag{
	// 	Name:  "h",
	// 	Value: "",
	// 	Usage: "Update Settings",
	// },
	// cli.StringFlag{
	// 	Name:  "config, c",
	// 	Usage: "Load configuration from `FILE`",
	// },
	}

	app.Commands = []cli.Command{
		{
			Name:    "update",
			Aliases: []string{"u"},
			Usage:   "Update home locale",
			Action: func(c *cli.Context) error {
				// fmt.Println("new task template: ", c.Args().First())
				if c.Args().First() == "home" {
					updateHomeLocale()
				}
				return nil
			},
		},
		{
			Name:    "get",
			Aliases: []string{"g"},
			Usage:   "Get weather by city",
			Action: func(c *cli.Context) error {
				cityReport(c.Args().First())
				return nil
			},
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	app.Action = func(c *cli.Context) error {

		//todo: if no other commands, check for home locale
		homeReport()
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

// func forecastByCoordinates(lat float64, lng float64) {
// 	darksky.Forecast(lat, lng)
// }

func homeReport() error {
	lat, lon := getHomeCoordinates()
	floatLat, err := strconv.ParseFloat(lat, 4)
	fatalError(err, "Invalid Latitude")

	floatLng, err := strconv.ParseFloat(lon, 4)
	fatalError(err, "Invalid Longitude")

	darksky.Forecast(appconfig.Config.Home.City, floatLat, floatLng)
	// forecastByCoordinates(floatLat, floatLon)
	return nil
}

func cityReport(city string) error {
	c, err := maps.NewClient(maps.WithAPIKey(""))
	fatalError(err, "Could not create request to get weather by city.")

	cty := searchForCity(city, c)
	floatLat, floatLng := getCityCoordinates(cty, c)
	darksky.Forecast(city, floatLat, floatLng)
	// forecastByCoordinates(floatLat, floatLng)

	return nil
}

func getHomeCoordinates() (string, string) {
	//If latitude, longitude, or home city are blank, update home locale
	if appconfig.Config.Home.Latitude == "" || appconfig.Config.Home.Longitude == "" || appconfig.Config.Home.City == "" {
		println("Home locale not set")
		updateHomeLocale()
	}
	return appconfig.Config.Home.Latitude, appconfig.Config.Home.Longitude
}

func updateHomeLocale() {
	c, err := maps.NewClient(maps.WithAPIKey(""))
	fatalError(err, "Could not create request to update home locale.")

	/*Get city in string form from citySearchPrompt() and send it searchForCity()
	which calls Google Places API to get list of autocompleted locations*/
	city := searchForCity(citySearchPrompt(), c)
	floatLat, floatLng := getCityCoordinates(city, c)
	lat := strconv.FormatFloat(floatLat, 'f', 4, 64)
	lng := strconv.FormatFloat(floatLng, 'f', 4, 64)
	appconfig.SetHomeLocale(city[0], city[1], lat, lng)
}

//Call to Google Geocode API to get Lat/Long from a city, state
func getCityCoordinates(locale []string, c *maps.Client) (float64, float64) {
	for i := range locale {
		strings.Replace(locale[i], " ", "+", -1)
	}
	address := strings.Join(locale, ",+")
	r := &maps.GeocodingRequest{
		Address: address,
	}
	result, err := c.Geocode(context.Background(), r)
	fatalError(err, "geocode search request")

	return result[0].Geometry.Location.Lat, result[0].Geometry.Location.Lng
}

//Get list of autocompleted location names from Google PlacesAPI
func searchForCity(city string, c *maps.Client) []string {
	r := &maps.PlaceAutocompleteRequest{
		Input: city,
		Types: "(cities)",
	}
	result, err := c.PlaceAutocomplete(context.Background(), r)
	if err != nil {
		log.Fatal("city search request failed2")
	}

	items := []string{}
	for _, p := range result.Predictions {
		items = append(items, p.Description)
	}

	prompt := promptui.Select{
		Label: "Select City",
		Items: items,
	}

	_, choice, err := prompt.Run()
	fatalError(err, "Prompt failed")

	cityState := strings.Split(choice, ", ")
	fmt.Printf("You chose %q\n", choice)

	return cityState
}

func citySearchPrompt() string {
	prompt := promptui.Prompt{
		Label: "Search by City:",
		// Validate: validate, //todo: validate string city
	}

	result, err := prompt.Run()
	fatalError(err, "Prompt failed")

	return result
}

func fatalError(err error, message string) {
	if err != nil {
		log.Fatalf(message+": %v\n", err)
	}
}
