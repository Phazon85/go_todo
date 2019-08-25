package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	configFile = "../dev.yaml"
)

func TestDBInit(t *testing.T) {
	db := DBInit(configFile)
	assert.NotNil(t, db)
}
