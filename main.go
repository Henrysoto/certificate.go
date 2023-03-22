package main

import (
	"fmt"
	"os"

	"certificate.go/cert"
	"certificate.go/pdf"
)

func main() {
	c, err := cert.New("Golang Programming", "John Doe", "2023-03-22")
	if err != nil {
		fmt.Printf("Error during certificate creation: %v", err)
		os.Exit(1)
	}
	var saver cert.Saver
	saver, err = pdf.New("output")
	if err != nil {
		fmt.Printf("Error during pdf creation: %v", err)
		os.Exit(1)
	}
	saver.Save(*c)
}
