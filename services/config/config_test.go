package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewConfig(t *testing.T) {
	cfg := NewConfig(configFile)
	assert.NotNil(t, cfg)
}

func TestLoadConfig(t *testing.T) {
	_ = NewConfig("")
}
