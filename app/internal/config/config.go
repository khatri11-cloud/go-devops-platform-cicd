package config

import "os"

type Config struct {
	Port       string
	AppVersion string
}

func Load() Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	version := os.Getenv("APP_VERSION")
	if version == "" {
		version = "v1.0.0"
	}

	return Config{
		Port:       port,
		AppVersion: version,
	}
}
