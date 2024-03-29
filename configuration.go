package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/deefstes/ClanEventsBot/logging"
	"github.com/deefstes/envtag"

	yaml "gopkg.in/yaml.v2"
)

// Configuration contains system wide configuration values
type Configuration struct {
	Token         string `yaml:"Token" ENV:"TOKEN"`
	CommandPrefix string `yaml:"CommandPrefix" ENV:"CMDPREFIX"`
	MongoDB       string `yaml:"MongoDB" ENV:"MONGODB"`
	ServiceTimer  int64  `yaml:"ServiceTimer" ENV:"SVCTIMER"`
	DebugLevel    int    `yaml:"DebugLevel" ENV:"LOGLEVEL"`
	HTTPPort      int    `yaml:"HTTPPort" ENV:"PORT"`
	APIKey        string `yaml:"APIKey" ENV:"APIKEY"`
}

// ReadConfigYAML reads system configuration from a YAML config file and returns a Configuration struct
func ReadConfigYAML() (Configuration, error) {
	var AppConfig Configuration
	exeFullPath, err := os.Executable()
	if err != nil {
		log.Println(logging.LogEntry{
			Severity: "ERROR",
			Message:  fmt.Sprintf("getting full executable path: %+v", err),
		})
		return AppConfig, err
	}

	ExeDirPath, err := filepath.Abs(filepath.Dir(exeFullPath))
	if err != nil {
		log.Println(logging.LogEntry{
			Severity: "ERROR",
			Message:  fmt.Sprintf("getting absolute executable path: %+v", err),
		})
		return AppConfig, err
	}

	yamlFile, err := ioutil.ReadFile(filepath.Join(ExeDirPath, "ClanEventsBot.yaml"))
	if err != nil {
		log.Println(logging.LogEntry{
			Severity: "ERROR",
			Message:  fmt.Sprintf("reading config file: %+v", err),
		})
		return AppConfig, err
	}

	err = yaml.Unmarshal(yamlFile, &AppConfig)
	if err != nil {
		log.Println(logging.LogEntry{
			Severity: "ERROR",
			Message:  fmt.Sprintf("unmarshalling config file: %+v", err),
		})
		return AppConfig, err
	}

	return AppConfig, nil
}

// ReadConfigENV reads system config from ENV variables
func ReadConfigENV() (Configuration, error) {
	var AppConfig Configuration
	err := envtag.Unmarshal("CEB_", &AppConfig)
	if err != nil {
		return AppConfig, fmt.Errorf("obtaining ENV value(s): %w", err)
	}

	return AppConfig, nil
}
