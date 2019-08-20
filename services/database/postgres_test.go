package database

import (
	"testing"
)

var config = "../../dev.yaml"

func TestDBInit(t *testing.T) {
	config := config.NewConfig(config)
	db := DBInit("postgres", config)
}
