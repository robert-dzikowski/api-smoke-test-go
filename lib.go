package main

import (
	"context"
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/robert-dzikowski/api-smoke-test-go/hrm"
)

func run(args argStruct) {
	// 2. Parse OAS spec from file or internet
	ctx := context.Background()
	loader := &openapi3.Loader{Context: ctx}

	// 2.1 Read and parse OAS file
	var doc *openapi3.T
	var err error

	if strings.HasPrefix(*args.oasFile, "http") {
		parsedURL, e := url.Parse(*args.oasFile)
		hrm.CheckError(e)
		doc, err = loader.LoadFromURI(parsedURL)
	} else {
		doc, err = loader.LoadFromFile(*args.oasFile) //("specs/petstore.json") //(*args.oasFile)
	}

	hrm.CheckError(err)

	// Validate OAS document
	if *args.validate {
		err = doc.Validate(ctx)
		hrm.CheckError(err)
	}

	baseApiUrl := doc.Servers[0].URL
	fmt.Println("Base URL:", baseApiUrl)
	fmt.Println("")
	fmt.Println("Testing", doc.Info.Title)
	fmt.Println("")

	// 3. Create list of GET endpoints
	endpointsList := getListOfParameterlessGETendpoints(doc)
	myLog(fmt.Sprint("Parmeterless GET endpoints: ", endpointsList))

	endpointsWithParams := getListOfGETendpointsWithParams(doc)
	myLog(fmt.Sprint("GET endpoints: ", endpointsWithParams))

	// 4. If endpointsList and endpointsWithParams are empty exit
	if len(endpointsList) == 0 && len(endpointsWithParams) == 0 {
		fmt.Println(
			"Test failed: spec file " + *args.oasFile + " doesn't contain any GET endpoints.")
		os.Exit(1)
	}

	// 5. Get auth token
	// token = None
	// if authorization_is_necessary():
	//     token = get_auth_token()

	// 6. Test parameterless GET endpoints
	hrm := hrm.New(baseApiUrl, "")
	fmt.Println("Testing GET methods")
	hrm.MakeGETRequests(endpointsList)

	// 7. Test GET endpoints that contain parameters
	// req_param = get_request_param_arg()
	//
	// if len(endpoints_with_params) > 0:
	//     call_get_methods_with_parameters(
	//         endpoints_with_params, maker, req_param)

	// 8. Print test results.
	// print_test_results(maker, spec['info']['title'])

	// 9. Exit with error if any test failed.
	// Exit with error code is needed by Azure to show test as failed
	// if len(maker.failed_requests_list) > 0:
	//     sys.exit(1)

	// 10. Exit with error if any test returned warning.
	// if config.WARNING_FAIL and len(maker.warning_requests_list) > 0:
	//     sys.exit(1)
}

func getListOfParameterlessGETendpoints(oasDoc *openapi3.T) []string {
	result := []string{}

	for path, pathItem := range oasDoc.Paths {
		for method := range pathItem.Operations() {
			if method == "GET" && !strings.Contains(path, "{") {
				result = append(result, path)
			}
		}
	}
	return result
}

func getListOfGETendpointsWithParams(oasDoc *openapi3.T) []string {
	result := []string{}

	for path, pathItem := range oasDoc.Paths {
		for method := range pathItem.Operations() {
			if method == "GET" && strings.Contains(path, "{") {
				result = append(result, path)
			}
		}
	}
	return result
}

func myLog(msg string) {
	const DEBUG bool = true
	if DEBUG {
		fmt.Println("log:", msg)
	}
}
