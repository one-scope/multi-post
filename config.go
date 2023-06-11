package main

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

var ServiceMap map[string]interface {
	SetCredentials(string)
}

type config struct {
	Services map[string]service   `yaml:"services"` //大文字じゃないとダメ
	Channels map[string][]channel `yaml:"channels"`
}

type service struct {
	Type        string `yaml:"type"`
	Credentials string `yaml:"credentials"`
}

type channel struct {
	Service string `yaml:"service"`
	Channel string `yaml:"channel"`
}

func ReadConfig(aFile string) (*config, error) {
	tConfig := config{}
	tByte, tError := os.ReadFile(aFile)
	if tError != nil {
		return nil, fmt.Errorf("failed to read %s: %w", aFile, tError)
	}
	if tError := yaml.Unmarshal(tByte, &tConfig); tError != nil {
		return nil, fmt.Errorf("failed to extract config from %s: %w", aFile, tError)
	}

	return &tConfig, nil
}
