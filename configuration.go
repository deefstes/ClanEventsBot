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
	MongoDB       string `yaml:"MongoDB`
}

func ReadConfig() (Configuration, error) {
	var AppConfig Configuration
	exeFullPath, err := os.Executable()
	if err != nil {
		fmt.Println("Error getting full executable path: %s", err.Error())
		return AppConfig, err
	}

	ExeDirPath, err := filepath.Abs(filepath.Dir(exeFullPath))
	if err != nil {
		fmt.Println("Error getting absolute executable path: %s", err.Error())
		return AppConfig, err
	}

	yamlFile, err := ioutil.ReadFile(filepath.Join(ExeDirPath, "ClanEventsBot.yaml"))
	if err != nil {
		fmt.Println("Error reading config file: %v", err)
		return AppConfig, err
	}

	err = yaml.Unmarshal(yamlFile, &AppConfig)
	if err != nil {
		fmt.Println("Error unmarshalling config file: %v", err)
		return AppConfig, err
	}

	return AppConfig, nil
}