package main

import (
	"fmt"
	"os"

	"github.com/sourcekris/idarc/lib/idarc"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("IDArc Go Port\n")
		fmt.Printf("Syntax: idarc archive_name\n")
		os.Exit(0)
	}

	path := os.Args[1]
	detector, err := idarc.NewDetector(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(int(idarc.FileNotFound))
	}

	result := detector.Detect()

	typeName := idarc.PackerNames[result.Type]
	if typeName == "" {
		typeName = "unknown"
	}

	suffix := ""
	if result.MultipleVolume && result.AV {
		suffix = "  (multiple volume archive + AV-secured/locked)"
	} else if result.MultipleVolume {
		suffix = "  (multiple volume archive)"
	} else if result.AV {
		suffix = "  (AV-secured/locked)"
	}

	fmt.Printf("Archive type = %s%s (%d)\n", typeName, suffix, result.Type)
	os.Exit(int(result.Type))
}
