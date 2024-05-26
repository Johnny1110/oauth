package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"oauth/sys"
	"sync"
)

var (
	sysProp *Properties
	once    sync.Once
)

func readConfig(filename string) Properties {
	// read yml
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		sys.Logger().Errorf("Error reading config file: %v", err)
	}

	// 解析 YAML
	var prop Properties
	if err := yaml.Unmarshal(data, &prop); err != nil {
		sys.Logger().Errorf("Error parsing config file: %v", err)
	}

	return prop
}

func GetProperties() *Properties {

	once.Do(func() {
		prop := readConfig("resources/profile.yml")

		switch prop.Profile {
		case "local":
			prop = readConfig("resources/local-config.yml")
			prop.Profile = "local"
			break
		case "prod":
			prop = readConfig("resources/prod-config.yml")
			prop.Profile = "prod"
			break
		default:
			panic("Invalid profile name.")
		}
		sys.Logger().Infof("active profile: " + prop.Profile)
		sysProp = &prop
	})
	return sysProp
}
