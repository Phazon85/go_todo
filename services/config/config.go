package config

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Config is the database connection info
type Config struct {
	Service *PostgresConfig `yaml: postgres`
}

// PostgresConfig contains the info for connecting to a database
type PostgresConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
	DB       string
}

// NewConfig loads the info from the config file
func NewConfig(file string) *Config {
	cfg := &Config{}
	if err := load(cfg, file); err != nil {
		panic(fmt.Sprintf("Error loading config from file: %s", err.Error()))
	}
	return cfg
}

// LoadValues takes the
func LoadValues(config interface{}, fname string) error {
	data, err := ioutil.ReadFile(fname)
	if err != nil {
		return err
	}

	if err := yaml.Unmarshal(data, config); err != nil {
		return err
	}
	return nil
}
