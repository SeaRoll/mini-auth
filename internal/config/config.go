package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Port string `yaml:"port"`
}

func NewConfig(path string) Config {
	c := Config{}
	contentByte, err := os.ReadFile(path)
	if err != nil {
		panic("Failed to open file")
	}

	err = yaml.Unmarshal([]byte(contentByte), &c)
	if err != nil {
		panic("Failed to unmarschal config")
	}

	return c
}