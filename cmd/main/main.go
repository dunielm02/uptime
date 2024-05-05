package main

import (
	"context"
	"flag"
	"lifeChecker/config"
	"lifeChecker/database"
	"lifeChecker/serviceSelector"
)

func main() {
	var configFile = flag.String("config", "uptime-config.yml", "Set the name of the configuration file.")
	flag.Parse()

	config := config.GetConfigFromYamlFile(*configFile)

	db := database.GetDatabaseFromConfig(config.Database)

	db.Connect()

	selector := serviceSelector.SelectorFromConfig(config.Services)

	selector.RunChecking(context.Background(), db)
}
