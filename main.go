package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	oasFile := flag.String(
		"oas", "", "Required, url or file name of the OpenAPI v.3 specification file")
	auth := flag.Bool("auth", false,
		"Use authentication, i.e. authentication token is used")
	localhost := flag.Bool("localhost", false,
		"Use when testing API that runs on your machine")
	onlyGet := flag.Bool("only-get", false, "Test only GET requests")
	requestParam := flag.Int("req-param", 13,
		"Value used in requests that contain parameters")
	help := flag.Bool("help", false, "Show help")

	flag.Parse()

	if *help {
		flag.PrintDefaults()
		os.Exit(0)
	}

	if *oasFile == "" {
		fmt.Println("\"oas\" argument is required")
		flag.PrintDefaults()
		os.Exit(1)
	}
	fmt.Println("oas:", *oasFile)

	if *auth {
		fmt.Println("auth:", *auth)
	}
	if *localhost {
		fmt.Println("localhost:", *localhost)
	}
	if *onlyGet {
		fmt.Println("only-get:", *onlyGet)
	}
	fmt.Println("req-param:", *requestParam)

}
