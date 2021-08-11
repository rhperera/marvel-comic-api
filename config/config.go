package config

import (
	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
)

var configMap map[string]string

func Init() {
	var err error
	configMap, err = godotenv.Read(".env")
	if err != nil {
		log.Fatal("Error reading configs ", err)
	}
	log.Info("Configs load success")
}

func Get(configName string) string {
	value := configMap[configName]
	if value == "" {
		log.Fatal("Invalid Config Read")
	}
	return value
}
