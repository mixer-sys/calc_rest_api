package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig_ValidFile(t *testing.T) {

	tempFile, err := os.CreateTemp("", "config.yaml")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tempFile.Name())
	configData := `
server:
  host: localhost
  port: "8080"
`
	if _, err := tempFile.Write([]byte(configData)); err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	tempFile.Close()

	config, err := LoadConfig(tempFile.Name())
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}

	assert.Equal(t, "localhost", config.Server.Host)
	assert.Equal(t, "8080", config.Server.Port)
}

func TestLoadConfig_InvalidFile(t *testing.T) {

	_, err := LoadConfig("invalidfile.yaml")
	assert.Error(t, err)
}
