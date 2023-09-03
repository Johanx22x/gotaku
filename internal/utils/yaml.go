package utils

import (
	"errors"
	"os"

	"gopkg.in/yaml.v2"
)

// DumpConfig dumps the preferences to a yaml file
func DumpConfig(preferences *Preferences) error {
    path, err := os.UserConfigDir()
    if err != nil {
        return errors.New("Error getting config directory")
    }

    cfgFile, err := os.Create(path + "/gotaku/config.yaml")
    if err != nil {
        return errors.New("Error creating config file")
    }
    defer cfgFile.Close()

    encoder := yaml.NewEncoder(cfgFile)
    err = encoder.Encode(preferences)
    if err != nil {
        return errors.New("Error writing to config file")
    }

    return nil
}
