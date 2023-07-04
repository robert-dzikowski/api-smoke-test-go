package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	oasFile := flag.String(
		"oas", "", "Url or file name of the OpenAPI v.3 specification file")
	help := flag.Bool("help", false, "Show help")
	flag.Parse()

	if *help {
		flag.PrintDefaults()
		os.Exit(0)
	}

	if *oasFile == "" {
		fmt.Println("oas argument is required")
		flag.PrintDefaults()
		os.Exit(1)
	}

	flag.Parse()
	fmt.Println(*oasFile)
}
