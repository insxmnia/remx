package config

type Configuration struct {
	Database struct {
		Driver       string `mapstructure:"driver"`
		EndpointType string `mapstructure:"endpoint-type"`
		Endpoint     string `mapstructure:"endpoint"`
		Token        string `mapstructure:"token"`
	} `mapstructure:"database"`
	Project struct {
		Types     []string                     `mapstructure:"types"`
		Variants  map[string][]string          `mapstructure:"variant"`
		Templates map[string]map[string]string `mapstructure:"template"`
	} `mapstructure:"project"`
}

type ConfigParams struct {
	Path       []string
	Name, Type string
}
