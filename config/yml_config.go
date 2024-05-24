package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

func readConfig(filename string) Properties {
	// read yml
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	// 解析 YAML
	var prop Properties
	if err := yaml.Unmarshal(data, &prop); err != nil {
		log.Fatalf("Error parsing config file: %v", err)
	}

	return prop
}

func GetProperties() Properties {
	return readConfig("config.yml")
}
