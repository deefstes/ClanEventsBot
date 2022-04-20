package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	yaml "gopkg.in/yaml.v2"
)

// Configuration contains system wide configuration values
type Configuration struct {
	Token         string `yaml:"Token"`
	CommandPrefix string `yaml:"CommandPrefix"`
	MongoDB       string `yaml:"MongoDB"`
	ServiceTimer  int64  `yaml:"ServiceTimer"`
	DebugLevel    int    `yaml:"DebugLevel"`
	HttpPort      int    `yaml:"HTTPPort"`
}

// ReadConfig reads system configuration from a YAML config file and returns a Configuration struct
func ReadConfig() (Configuration, error) {
	var AppConfig Configuration
	exeFullPath, err := os.Executable()
	if err != nil {
		fmt.Printf("Error getting full executable path: %s", err.Error())
		return AppConfig, err
	}

	ExeDirPath, err := filepath.Abs(filepath.Dir(exeFullPath))
	if err != nil {
		fmt.Printf("Error getting absolute executable path: %s", err.Error())
		return AppConfig, err
	}

	yamlFile, err := ioutil.ReadFile(filepath.Join(ExeDirPath, "ClanEventsBot.yaml"))
	if err != nil {
		fmt.Printf("Error reading config file: %s", err.Error())
		return AppConfig, err
	}

	err = yaml.Unmarshal(yamlFile, &AppConfig)
	if err != nil {
		fmt.Printf("Error unmarshalling config file: %s", err.Error())
		return AppConfig, err
	}

	return AppConfig, nil
}
