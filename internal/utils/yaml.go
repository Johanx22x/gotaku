package utils

import (
    "os"

    "gopkg.in/yaml.v2"
)

// createDefaultConfig creates the default config file
func createDefaultConfig(path string, preferences *Preferences) error {
    cfgFile, err := os.Create(path + "/config.yaml")
    if err != nil {
        return err
    }
    defer cfgFile.Close()

    encoder := yaml.NewEncoder(cfgFile)
    err = encoder.Encode(preferences)
    if err != nil {
        return err
    }

    return nil
}
