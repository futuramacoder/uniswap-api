package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Environment struct {
	AppName     string   `envconfig:"APP_NAME" default:"uniswap-api"`
	Port        int      `envconfig:"PORT" default:"1337"`
	NodeUrl     string   `envconfig:"NODE_URL" required:"true"`
	LogLevel    string   `envconfig:"LOG_LEVEL" default:"trace"`
	CorsOrigins []string `envconfig:"CORS_ORIGINS" default:"*"`
	CorsMethods []string `envconfig:"CORS_METHODS"`
}

func LoadEnvConfig() (*Environment, error) {
	e := &Environment{}
	if err := envconfig.Process("", e); err != nil {
		return nil, err
	}

	return e, nil
}
