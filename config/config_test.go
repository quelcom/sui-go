package config

import (
	"testing"
)

func TestReadBrokenConfig(t *testing.T) {
	var conf = []byte(`abc`)
	_, err := setConfig(conf)
	if err == nil {
		t.Error("Expected error but got nil")
	}
}

func TestReadConfigNoPort(t *testing.T) {
	var conf = []byte(`name = "Gopher"
	port = 0`)
	_, err := setConfig(conf)
	if err == nil {
		t.Error("Expected error but got nil")
	}

	conf = []byte(`name = "Gopher"`)
	_, err = setConfig(conf)
	if err == nil {
		t.Error("Expected error but got nil")
	}
}

func TestReadConfigFromFile(t *testing.T) {
	c, err := ParseFromFile("config_test.toml")
	if err != nil {
		t.Errorf("Did not expect error but got %v", err)
	}

	if c.Name != "Gopher" {
		t.Errorf("Expected Gopher but got %s", c.Name)
	}

	if c.Providers[1].Name != "Discogs" {
		t.Errorf("Expected 'Discogs' but got %q", c.Providers[1].Name)
	}
}

func TestReadConfigFromFileDoesNotExist(t *testing.T) {
	_, err := ParseFromFile("config_test_does_not_exist.toml")
	if err == nil {
		t.Error("Expected error but got nil")
	}
}
