package config

type Properties struct {
	Profile string `yaml:"profile"`

	Port string `yaml:"port"`

	DB struct {
		ConnectionString      string `yaml:"connection-string"`
		IP                    string `yaml:"ip"`
		Port                  string `yaml:"port"`
		Username              string `yaml:"username"`
		Password              string `yaml:"password"`
		MaxConnection         int    `yaml:"max-connection"`
		MaxIdleConnection     int    `yaml:"max-idle-connection"`
		MaxConnectionLifetime int    `yaml:"max-connection-lifetime"` // in minutes
	} `yaml:"db"`

	Redis struct {
		IP       string `yaml:"ip"`
		Port     int    `yaml:"port"`
		Password string `yaml:"password"`
	} `yaml:"redis"`
}
