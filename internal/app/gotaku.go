package app

// App is the main application
type Gotaku struct {
    Config *Config
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
    }
}
