package main

import (
	"log"
	"os"
)

func main() {
	f, err := os.Open("datei-ist-nicht-da.txt")
	if err != nil {
		log.Fatalf("Konnte die Datei nicht öffnen: %v", err)
	}
	f.Close()
}
