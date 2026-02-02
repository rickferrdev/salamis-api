package config

import (
	"github.com/rickferrdev/dotenv"
)

type Env struct {
	AppPort      string `env:"APP_PORT"`
	AppMode      string `env:"APP_MODE"`
	AppSecretJWT string `env:"APP_SECRET_JWT"`
	AppDBUrl     string `env:"APP_DB_URL"`
}

func NewEnv() (*Env, error) {
	var env Env

	// Load .env file variables into the environment
	dotenv.Collect()

	if err := dotenv.Unmarshal(&env); err != nil {
		return nil, err
	}

	return &env, nil
}
