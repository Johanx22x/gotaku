package utils

import (
    "fmt"
    "os"
    "errors"

    "gopkg.in/yaml.v2"
)

type Preferences struct {
    Anime    string  `yaml:"anime"`
    Manga    string  `yaml:"manga"`
    Manager  string  `yaml:"manager"`
    Token    string  `yaml:"token"`
}

// defaultPreferences returns the default preferences
func defaultPreferences(path string) (*Preferences, error) {
    preferences := &Preferences{
        Anime: "9anime",
        Manga: "mangadex",
        Manager: "anilist",
        Token: "",
    }

    err := createDefaultConfig(path, preferences)
    if err != nil {
        return nil, errors.New("Error creating config file")
    }

    return preferences, nil
}

// LoadPreferences parses the config file and returns the user preferences
func LoadPreferences(verbose bool) (*Preferences, error) {
    cfgPath, err := os.UserConfigDir()
    if err != nil {
        return nil, errors.New("Error getting config directory")
    }

    // check if gotaku config directory exists
    cfgDir := cfgPath + "/gotaku"
    if _, err := os.Stat(cfgDir); os.IsNotExist(err) {
        err = os.Mkdir(cfgDir, 0755)
        if err != nil {
            return nil, errors.New("Error creating config directory")
        }
    }

    // check if config file exists
    cfgFile, err := os.Open(cfgDir + "/config.yaml")
    if err != nil {
        if verbose {
            fmt.Println("Config file not found, using default preferences")
        }
        return defaultPreferences(cfgDir)
    }
    defer cfgFile.Close()

    cfg := &Preferences{}

    decoder := yaml.NewDecoder(cfgFile)
    err = decoder.Decode(cfg)
    if err != nil {
        return nil, errors.New("Error parsing config file")
    }

    return cfg, nil
}
