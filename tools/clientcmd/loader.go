package clientcmd

import (
	"gopkg.in/yaml.v3"
	"os"
)

// LoadFormFile loads configuration from a file.
func LoadFormFile(filename string) (*Config, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	config, err := Load(data)
	if err != nil {
		return nil, err
	}
	config.AuthInfo.LocationOfOrigin = filename
	config.Server.LocationOfOrigin = filename
	if config.AuthInfo == nil {
		config.AuthInfo = &AuthInfo{}
	}
	if config.Server == nil {
		config.Server = &Server{}
	}
	return config, nil
}

// Load loads configuration from data.
func Load(data []byte) (*Config, error) {
	config := NewConfig()
	if len(data) == 0 {
		return config, nil
	}
	if err := yaml.Unmarshal(data, config); err != nil {
		return nil, err
	}
	return config, nil
}
