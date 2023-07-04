package main

import (
	"flag"
	"fmt"
	"os"
)

type argStruct struct {
	oasFile      *string
	auth         *bool
	localhost    *bool
	onlyGet      *bool
	requestParam *int
}

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

	args := argStruct{
		oasFile,
		auth,
		localhost,
		onlyGet,
		requestParam,
	}
	if *args.oasFile == "" {
		fmt.Println("\"oas\" argument is required")
		flag.PrintDefaults()
		os.Exit(1)
	}

	printArguments(args)

}

func printArguments(args argStruct) {
	fmt.Println("oas:", *args.oasFile)
	if *args.auth {
		fmt.Println("auth:", *args.auth)
	}
	if *args.localhost {
		fmt.Println("localhost:", *args.localhost)
	}
	if *args.onlyGet {
		fmt.Println("only-get:", *args.onlyGet)
	}
	fmt.Println("req-param:", *args.requestParam)
}
