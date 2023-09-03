package main

import (
    "fmt"

    "github.com/Johanx22x/gotaku"
)

func main() {
    config := gotaku.GetConfig()
    fmt.Printf("Gotaku v%s\n", config.Version)
}
