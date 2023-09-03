package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Johanx22x/gotaku/internal/app"
)

func main() {
    app := app.GetApp()

    // parse command line flags
    version := flag.Bool("v", false, "Prints the version of Gotaku")
    verbose := flag.Bool("verbose", false, "Prints verbose output")
    flag.Parse()

    if *version {
        fmt.Printf("Gotaku version %s\n", app.Config.Version)
        os.Exit(0)
    }

    if *verbose {
        fmt.Println("Verbose output enabled")
    }

    app.Run()
}
