package app

import (
	"fmt"
	"os"

	"github.com/Johanx22x/gotaku/internal/utils"
)

// App is the main application
type Gotaku struct {
    Config      *Config
    Preferences *utils.Preferences
}

// singleton instance of the app
var instance *Gotaku

// GetApp returns the singleton instance of the app 
func GetApp() *Gotaku {
    if instance == nil {
        instance = &Gotaku{
            Config: GetConfig(),
        }
    }

    return instance
}

// Run runs the application
func (app *Gotaku) Run() {
    if app.Config.Verbose {
        fmt.Printf("Running Gotaku v%s!\n", app.Config.Version)
    }

    preferences, err := utils.LoadPreferences(app.Config.Verbose)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    app.Preferences = preferences
    if app.Config.Verbose {
        fmt.Println("Preferences loaded!")
    }

    if app.Preferences.Token == "" {
        token, err := app.Authenticate()
        if err != nil {
            fmt.Println(err)
            os.Exit(1)
        }

        app.Preferences.Token = token
    }
}
