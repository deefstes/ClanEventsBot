package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	yaml "gopkg.in/yaml.v2"
)

type Configuration struct {
	Token         string `yaml:"Token"`
	CommandPrefix string `yaml:"CommandPrefix"`
}

func ReadConfig() (Configuration, bool) {
	var AppConfig Configuration
	exeFullPath, err := os.Executable()
	if err != nil {
		fmt.Println("Error getting full executable path: %s", err.Error())
		return AppConfig, false
	}

	ExeDirPath, err := filepath.Abs(filepath.Dir(exeFullPath))
	if err != nil {
		fmt.Println("Error getting absolute executable path: %s", err.Error())
		return AppConfig, false
	}

	yamlFile, err := ioutil.ReadFile(filepath.Join(ExeDirPath, "ClanEventsBot.yaml"))
	if err != nil {
		fmt.Println("Error reading config file: %v", err)
		return AppConfig, false
	}

	err = yaml.Unmarshal(yamlFile, &AppConfig)
	if err != nil {
		fmt.Println("Error unmarshalling config file: %v", err)
		return AppConfig, false
	}

	return AppConfig, true
}
