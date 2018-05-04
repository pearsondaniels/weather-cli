# weather-cli
Weather report command line application written in Go

## Installation Instructions
Clone the repository  
Create a configuration file named `config.json` inside the `appconfig` directory.  
This is what goes inside it:
```
{
    "home":{
        "city":"",
        "state":"",
        "latitude":"",
        "longitude":""
    }
}
```

Make sure you have your GOPATH and GOBIN set up  
Run `go install weather.go` inside the main directory

## Basic Usage
**NOTE: Current state of application requires that you run all comands inside the install directory. Fix upcoming.**  
#### Gets weather report for home locale
```
$ weather
```
#### Updates home locale
```
$ weather update home
```