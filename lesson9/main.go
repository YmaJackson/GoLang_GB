package main

import (
	"flag"
	"fmt"
	"GO_GB/lesson9/configuration"
	"os"
)

var jsonConfigPath string
var yamlCongigPath string

func main() {

	flag.StringVar(&jsonConfigPath, "json_config_path", "./configuration/conf.json", "Path for JSON config file")
	flag.StringVar(&yamlCongigPath, "yaml_config_path", "./configuration/conf.yaml", "Path for YAML config file")

	conf, err := configuration.InitConfig("env", "")
	if err != nil {
		fmt.Printf("Error loading the configuration: %s\n", err)
		os.Exit(1)
	}
	conf.PrintConfig("ENV")

	conf, err = configuration.InitConfig("json", jsonConfigPath)
	if err != nil {
		fmt.Printf("Error loading the configuration: %s\n", err)
		os.Exit(1)
	}
	conf.PrintConfig("JSON")

	conf, err = configuration.InitConfig("yaml", yamlCongigPath)
	if err != nil {
		fmt.Printf("Error loading the configuration: %s\n", err)
		os.Exit(1)
	}
	conf.PrintConfig("YAML")

	fmt.Printf("Configuration loading \n")
}
