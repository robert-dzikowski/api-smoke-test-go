package main

import (
	"flag"
	"fmt"
	"os"
)

type argStruct struct {
	// Url or file name of the OpenAPI v.3 specification file
	oasFile      *string
	auth         *bool
	localhost    *bool
	onlyGet      *bool
	requestParam *int
	validate     *bool
	singleThread *bool
}

func main() {
	oasFile := flag.String(
		"oas", "", "Required, file name of the OpenAPI 3.0 specification file")

	auth := flag.Bool("auth", false,
		"Use authentication, i.e. authentication token is used to authorize requests")

	localhost := flag.Bool("localhost", false,
		"Not implemented. Use when testing API that runs on your local machine")

	onlyGet := flag.Bool("only-get", false, "Not implemented. Test only GET requests")

	requestParam := flag.Int("req-param", 13,
		"Integer used in requests that contain parameters")

	validate := flag.Bool("validate", false, "Validate file given to \"oas\" argument")

	singleThread := flag.Bool("single-thread", false,
		"Use single thread for HTTP requests.")

	help := flag.Bool("help", false, "Show help")

	flag.Parse()

	if *help {
		flag.PrintDefaults()
		os.Exit(0)
	}

	//*oasFile = "specs/api.github.com.json"
	if *oasFile == "" {
		fmt.Println("\"oas\" argument is required")
		flag.PrintDefaults()
		os.Exit(1)
	}

	args := argStruct{
		oasFile,
		auth,
		localhost,
		onlyGet,
		requestParam,
		validate,
		singleThread,
	}
	printArguments(args)

	//printToken()

	run(args)
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
	if *args.validate {
		fmt.Println("validate:", *args.validate)
	}
}
