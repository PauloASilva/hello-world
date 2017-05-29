package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

// Application configurations
type Config struct {
	Port int
}

// Loads configurations from configFile into conf struct
func (conf *Config) Load(configFile string) *Config {
	yamlFile, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Fatalf("[ERROR] %v", err)
	}

	err = yaml.Unmarshal(yamlFile, conf)
	if err != nil {
		log.Fatalf("[ERROR] %v", err)
	}

	return conf
}
