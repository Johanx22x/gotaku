package app

import (
	"fmt"
	"os"

	"github.com/Johanx22x/gotaku/internal/api/usermanager"
    "github.com/Johanx22x/gotaku/internal/api/anime"
    "github.com/Johanx22x/gotaku/internal/api/manga"
    "github.com/Johanx22x/gotaku/internal/utils"
)

// App is the main application
type Gotaku struct {
    Config      *Config
    Preferences *utils.Preferences
    UserManager usermanager.API
    Anime       anime.API 
    Manga       manga.API
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

    app.UserManager, err = usermanager.GetManager(app.Preferences.Manager)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    if app.Preferences.Token == "" {
        if app.Config.Verbose {
            fmt.Println("No cached token found, authenticating...")
        }
        token, err := app.UserManager.Authenticate()
        if err != nil {
            fmt.Println(err)
            os.Exit(1)
        }

        app.Preferences.Token = token
        err = utils.DumpConfig(app.Preferences)
        if err != nil {
            fmt.Println(err)
            os.Exit(1)
        }
    } else {
        if app.Config.Verbose {
            fmt.Println("Using cached token!")
        }
    }
}
