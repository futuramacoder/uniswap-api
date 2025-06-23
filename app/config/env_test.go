package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadEnvConfig_DefaultsAndRequired(t *testing.T) {
	t.Setenv("PORT", "1337")
	t.Setenv("NODE_URL", "https://rpc.testnet")
	t.Setenv("LOG_LEVEL", "trace")
	t.Setenv("CORS_ORIGINS", "*")
	t.Setenv("CORS_METHODS", "GET,POST")

	cfg, err := LoadEnvConfig()
	assert.NoError(t, err)
	assert.NotNil(t, cfg)

	assert.Equal(t, 1337, cfg.Port)
	assert.Equal(t, "https://rpc.testnet", cfg.NodeUrl)
	assert.Equal(t, "trace", cfg.LogLevel)
	assert.Equal(t, []string{"*"}, cfg.CorsOrigins)
	assert.Equal(t, []string{"GET", "POST"}, cfg.CorsMethods)
}

func TestLoadEnvConfig_MissingRequired(t *testing.T) {
	_, err := LoadEnvConfig()
	assert.Error(t, err)
}
