package postgres

import (
	"testing"

	"github.com/phazon85/go_todo/services/config"
	"github.com/stretchr/testify/assert"
)

const (
	configFile = "../../dev.yaml"
)

func TestDBInit(t *testing.T) {
	config := config.NewConfig(configFile)
	db := DBInit(config)
	assert.NotNil(t, db)
}
