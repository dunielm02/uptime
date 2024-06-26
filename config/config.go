package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type DatabaseConfig struct {
	Type string `yaml:"type"`
	Spec any    `yaml:"spec"`
}

type ServiceConfig struct {
	Type                 string   `yaml:"type"`
	Name                 string   `yaml:"name"`
	WaitingTime          int      `yaml:"waiting-time"`
	Timeout              int      `yaml:"timeout"`
	Inverted             bool     `yaml:"inverted"`
	NotificationChannels []string `yaml:"notification-channels"`
	Spec                 any      `yaml:"spec"`
}

type PortForwardConfig struct {
	ServerAddress string `yaml:"server-address"`
	RemoteAddress string `yaml:"remote-address"`
	LocalAddress  string `yaml:"local-address"`
	Username      string `yaml:"username"`
	Password      string `yaml:"password"`
}

type NotificationChannelsConfig struct {
	Name string `yaml:"name"`
	Type string `yaml:"type"`
	Spec any
}

type Config struct {
	Database             DatabaseConfig               `yaml:"database"`
	Services             []ServiceConfig              `yaml:"services"`
	PortForwards         []PortForwardConfig          `yaml:"port-forward"`
	NotificationChannels []NotificationChannelsConfig `yaml:"notification-channels"`
}

func GetConfigFromYamlFile(fileName string) Config {
	yamlData, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal("Error reading the config file: ", err)
	}

	var config Config

	err = yaml.Unmarshal(yamlData, &config)
	if err != nil {
		log.Fatal(err)
	}

	return config
}
