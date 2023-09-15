package main

import (
	"log"
	"os"

	yaml "gopkg.in/yaml.v3"
)

type CFG struct {
	// Server Stuff
	URL         string `yaml:"url"`
	Ip          string `yaml:"ip"`
	ServicePort int    `yaml:"port"`

	// Database
	Host             string `yaml:"host"`
	DatabasePort     int    `yaml:"port"`
	DatabaseUser     string `yaml:"user"`
	DatabasePassword string `yaml:"password"`
	DBName           string `yaml:"dbname"`
	SSLMode          bool   `yaml:"sslmode"`
	TimeZone         string `yaml:"timezone"`

	// Initial Admin
	Enabled bool   `yaml:"enabled"`
	Token   string `yaml:"token"`
	Admin   bool   `yaml:"admin"`
}

func GetConfig() *CFG {
	return parseConfig()
}

func parseConfig() *CFG {
	yFile, err := os.ReadFile("config.yml")
	if err != nil {
		log.SetFlags(log.Lshortfile | log.LstdFlags)
		log.Fatalf("Error opening yml config file: %v\n", err)
	}

	var cfg CFG

	err = yaml.Unmarshal(yFile, &cfg)
	if err != nil {
		log.SetFlags(log.Lshortfile | log.LstdFlags)
		log.Fatalf("Error unmarshaling yaml config file: %v\n", err)
	}

	return &cfg
}
