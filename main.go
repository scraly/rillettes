package main

import (
	"fmt"
	"log"

	"embed"
)

// Hey, I want to embed "rillettes" folder in the executable binary
// Use embed go 1.16 new feature (for embed gophers static files)
//
//go:embed rillettes
var embedRillettesFiles embed.FS

func main() {

	// Display rillettes asii embed files
	fileData, err := embedRillettesFiles.ReadFile("rillettes/rillettes.txt")
	if err != nil {
		log.Fatal("Error during read rillettes ascii file", err)
	}
	fmt.Println(string(fileData))

}
