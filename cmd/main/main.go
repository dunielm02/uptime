package main

import (
	"context"
	"flag"
	"lifeChecker/config"
	"lifeChecker/database"
	"lifeChecker/serviceSelector"
	"log"
)

func main() {
	var configFile = flag.String("config", "uptime-config.yml", "Set the name of the configuration file.")
	flag.Parse()

	config := config.GetConfigFromYamlFile(*configFile)

	db := database.GetDatabaseFromConfig(config.Database)

	err := db.Connect()
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}

	selector := serviceSelector.SelectorFromConfig(config)

	selector.RunChecking(context.Background(), db)
}
