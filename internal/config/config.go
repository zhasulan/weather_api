package config

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"weather_api/internal/database"
)

type App struct {
	Name          string `json:"name"`
	Host          string `json:"host"` // can be empty
	Port          string `json:"port"`
	IsDevelopment bool   `json:"is_development"`
	LogLevel      string `json:"log_level"`
	Version       string `json:"version"`
}

type WeatherApiConfig struct {
	HostURL string `json:"host_url"`
	Key     string `json:"key"`
	Timeout int    `json:"timeout"`
}

type Configuration struct {
	App        App              `json:"app"`
	DB         database.Config  `json:"db"`
	WeatherApi WeatherApiConfig `json:"weather_api"`
}

var Config *Configuration // not init with = &Configuration{}, we wait panic if file not found or another error on initialization of config

func GetConfigPath() string {
	projectDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Can't find work directory: %s", err.Error())
	}
	var path string
	flag.StringVar(&path, "config", fmt.Sprintf("%s/config/conf.json", projectDir), "Config path")
	flag.Parse()
	return path
}

func InitConfig(path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Can't open config file: %s %v ", path, err)
	}

	defer func(file *os.File) {
		if err = file.Close(); err != nil {
			log.Fatalf("Can't close config file: %s %v ", path, err)
		}
	}(file)

	decoder := json.NewDecoder(file)
	if err = decoder.Decode(&Config); err != nil {
		log.Fatalf("Decoding config error: %v", err)
	}
}
