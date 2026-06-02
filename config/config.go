package config

import "os"

type Config struct {
	URL   string
	Token string
}

// Load loads configuration from environment variables
func Load() *Config {
	url := os.Getenv("URL")
	Token := os.Getenv("TOKEN")
	return &Config{
		URL:   url,
		Token: Token,
	}
}
