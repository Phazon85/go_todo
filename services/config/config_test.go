package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var configFile = "../../dev.yaml"

func TestNewConfig(t *testing.T) {
	cfg := NewConfig(configFile)
	assert.NotNil(t, cfg)
}

// func TestLoadConfig(t *testing.T) {
// 	_ = NewConfig("")
// }
