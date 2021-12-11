package main

import (
	"fmt"
	"gb_level1/lesson8/configuration"
	"os"
)

func main() {
	conf, err := configuration.InitConfig()
	if err != nil {
		fmt.Printf("Error load configuration\n")
		os.Exit(1)
	}

	fmt.Printf("Configuration from ENV: %+v\n\n", conf)