package config

import "github.com/kelseyhightower/envconfig"

type (
	Specification struct {
		Database DatabaseSpecification
	}

	DatabaseSpecification struct {
		Host     string `envconfig:"HOST"`
		Port     string `envconfig:"PORT"`
		User     string `envconfig:"USER"`
		Password string `envconfig:"PASSWORD"`
		Dbname   string `envconfig:"DBNAME"`
		Sslmode  string `envconfig:"SSLMODE"`
	}
)

var Settings Specification

func Init() {
	if err := envconfig.Process("", &Settings); err != nil {
		panic(err)
	}
}
