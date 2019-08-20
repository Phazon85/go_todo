package config

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Config is the database connection info
type Config struct {
	Service *PostgresConfig `yaml:"postgres"`
}

// PostgresConfig contains the info for connecting to a database
type PostgresConfig struct {
	Host     string `yaml:"host"`
	Port     int64  `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
	DB       string `yaml:"db"`
}

// NewConfig loads the info from the config file
func NewConfig(file string) *Config {
	cfg := &Config{}
	if err := load(cfg, file); err != nil {
		panic(fmt.Sprintf("failed to load config file %s", err.Error()))
	}
	return cfg
}

// LoadValues takes the
func load(config interface{}, fname string) error {
	data, err := ioutil.ReadFile(fname)
	if err != nil {
		return err
	}

	if err := yaml.Unmarshal(data, config); err != nil {
		return err
	}
	return nil
}
