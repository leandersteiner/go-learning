package main

import (
	"archive/zip"
	"bytes"
	"fmt"
)

func main() {
	data := []byte("Not a ZIP file")
	notAZipFile := bytes.NewReader(data)
	_, err := zip.NewReader(notAZipFile, int64(len(data)))
	if err == zip.ErrFormat {
		fmt.Println("Told you so")
	}
}
