package main

import "github.com/caarlos0/env/v9"

type Config struct {
	Auth struct {
		Issuer   string `env:"OIDC_ISSUER"`
		ClientID string `env:"OIDC_CLIENT_ID"`
		KeyID    string `env:"OIDC_KEY_ID"`
		Key      string `env:"OIDC_KEY"`
	}
}

func ConfigFromEnv() (Config, error) {
	cnf := Config{}
	err := env.Parse(&cnf)
	return cnf, err
}
