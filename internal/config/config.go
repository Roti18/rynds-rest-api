package config

import "os"

type Config struct {
	Appname string
}

func Load() Config {
	return Config{
		Appname: getEnv("APP_NAME", "Rynds Api"),
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
