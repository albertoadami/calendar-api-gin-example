package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	config, err := LoadConfig()

	assert.Equal(t, nil, err)
	assert.Equal(t, "localhost", config.DatabaseHost)
	assert.Equal(t, 5432, config.DatabasePort)
	assert.Equal(t, "postgres", config.DatabaseUser)
	assert.Equal(t, "postgres", config.DatabasePassword)
	assert.Equal(t, "calendar-api", config.DatabaseName)
}
