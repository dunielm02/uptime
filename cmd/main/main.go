package main

import (
	"flag"
	"fmt"
)

func main() {
	var configFile = flag.String("config", "uptime-config.yml", "Set the name of the configuration file.")
	flag.Parse()

	fmt.Println(configFile)
}
