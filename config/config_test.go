package config

import (
	"testing"
)

func TestGetConfig(t *testing.T) {
	config := LoadConfig("../config.json")
	if config.Host == "" {
		t.Error("No config")
	}
	if config.Key == "" {
		t.Error("No config")
	}
	if config.Secret == "" {
		t.Error("No config")
	}
}

func TestGetMasterConfig(t *testing.T) {
	config := LoadMasterConfig("../config.json")
	if config.Master.Host == "" {
		t.Error("No config")
	}
}
