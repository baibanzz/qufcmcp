package config

import "os"

type Config struct {
	URL       string `yaml:"url"`
	MFASecret string `yaml:"mfa_secret"`
}

// Load loads configuration from environment variables
func Load() *Config {
	url := os.Getenv("URL")
	mfaSecret := os.Getenv("MFA_SECRET")
	return &Config{
		URL:       url,
		MFASecret: mfaSecret,
	}
}
