package main

import (
    "fmt"
    "os"
    "go-monkey-go/repl"
)

func main() {
    // Let's get ready to RUMBLE!
    fmt.Println("Hello! Welcome to MonkeyLang")
    repl.Start(os.Stdin, os.Stdout)
}
