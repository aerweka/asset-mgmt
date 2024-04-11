package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Database struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Name     string `yaml:"name"`
	}
	Cloudinary struct {
		CloudName string `yaml:"cloud_name"`
		APIKey    string `yaml:"api_key"`
		APISecret string `yaml:"api_secret"`
	}
}

func Init() (*Config, error) {
	c := Config{}

	cfgFile, err := os.ReadFile("config/config.yaml")
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(cfgFile, &c)
	if err != nil {
		return nil, err
	}

	return &c, nil
}
