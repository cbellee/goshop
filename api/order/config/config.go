// Package config contain configuration of the project.
// It's a good idea for configuration to have its own package to prevent the circular import because it naturally will be used by multiple packages.
// Don't forget to create test for each configuration field to make sure no typo and get the expected value/
package config

import (
	"github.com/kelseyhightower/envconfig"
)

// Config type
// Check https://github.com/kelseyhightower/envconfig#struct-tag-support for more help
type Config struct {
	Address      string `envconfig:"ADDRESS" default:"0.0.0.0"`
	Port         string `envconfig:"PORT" default:"80"`
	DbName       string `envconfig:"DB_NAME" required:"true"`
	DbHostName   string `envconfig:"DB_HOST_NAME" required:"true"`
	DbHostNameSuffix string `envconfig:"DB_HOST_NAME_SUFFIX" default:"documents.azure.com"`
	DbPassword   string `envconfig:"DB_PASSWORD" required:"true"`
	DbPort       int    `envconfig:"DB_PORT" default:"10255"`
	DbCollection string `envconfig:"DB_COLL" default:"orders"`
	IsLocal      bool   `envconfig:"DB_IS_LOCAL" default:"false"`
}

// LoadConfig return new instance of config
func LoadConfig() (conf Config, err error) {
	err = envconfig.Process(Prefix, &conf)
	return
}

// LoadConfigForTest return config for testing environment
func LoadConfigForTest() (conf Config, err error) {
	conf, err = LoadConfig()
	conf.DbName = conf.DbName + "_test"
	return
}

// Details return details of configuration
func Details() []ConfigDetail {
	return details(&Config{})
}
