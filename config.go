package gotaku

import (
)

// Config is a struct that holds the configuration for the application.
type Config struct {
    Version string      // Version of the application
}

// GetConfig returns a Config struct with the default values.
func GetConfig() Config {
    return Config{
        Version: "0.1.0",
    }
}
