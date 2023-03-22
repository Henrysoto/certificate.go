package main

import (
	"flag"
	"fmt"
	"os"

	"certificate.go/cert"
	"certificate.go/html"
	"certificate.go/pdf"
)

func main() {
	outputType := flag.String("type", "pdf", "Output type of the certificate.")
	fileName := flag.String("file", "", "Path to CSV file")
	flag.Parse()

	if len(*fileName) <= 0 {
		fmt.Printf("Error trying to read file.\n")
		os.Exit(1)
	}

	var saver cert.Saver
	var err error
	switch *outputType {
	case "html":
		saver, err = html.New("output")
	case "pdf":
		saver, err = pdf.New("output")
	default:
		fmt.Printf("Unknown output type: '%v'\n", outputType)
		os.Exit(1)
	}
	if err != nil {
		fmt.Printf("Error could not generate: %v\n", err)
		os.Exit(1)
	}

	certs, err := cert.ParseCSV(*fileName)
	if err != nil {
		fmt.Printf("Error could not parse csv file: %v\n", err)
		os.Exit(1)
	}

	for _, c := range certs {
		err = saver.Save(*c)
		if err != nil {
			fmt.Printf("Error trying to save certificate: %v\n", err)
		}
	}
}
