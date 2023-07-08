package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/robert-dzikowski/api-smoke-test-go/hrm"
)

const AUTH_TOKEN = "auth_token"

// File containing correct status codes for HTTP GET method
const SC_FLE = "status_codes.txt"

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
		doc, err = loader.LoadFromFile(*args.oasFile)
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
	token := ""
	if *args.auth {
		t, ok := os.LookupEnv(AUTH_TOKEN)
		if !ok {
			fmt.Printf("Error: env. variable %s is not set.", AUTH_TOKEN)
		}
		token = t
	}

	// 6. Test parameterless GET endpoints
	// GET corect status codes from file
	lines := getFileLines(SC_FLE)
	sc := getCorrectGETSC(lines)

	hrm := hrm.New(baseApiUrl, token, sc)
	fmt.Println("Testing GET methods")
	hrm.MakeGETRequests(endpointsList)

	// 7. Test GET endpoints that contain parameters
	// req_param = get_request_param_arg()
	//
	// if len(endpoints_with_params) > 0:
	//     call_get_methods_with_parameters(
	//         endpoints_with_params, maker, req_param)

	// 8. Print test results.
	printAndSaveTestResults(hrm, doc.Info.Title)

	// 9. Exit with error if any test failed.
	// Exit with error code is needed by Azure to show test as failed
	// if len(maker.failed_requests_list) > 0:
	//     sys.exit(1)

	// 10. Exit with error if any test returned warning.
	// if config.WARNING_FAIL and len(maker.warning_requests_list) > 0:
	//     sys.exit(1)
} // run()

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

func getFileLines(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func getCorrectGETSC(lines []string) []int {
	const GET_SC string = "GET Status Codes"
	scLine := ""

	for _, line := range lines {
		if strings.HasPrefix(line, GET_SC) {
			scLine = line
			break
		}
	}
	if len(scLine) == 0 {
		log.Fatalf("String '%s' was not found", GET_SC)
	}

	index := strings.Index(scLine, ":")
	if index == -1 {
		log.Fatalf("':' not found in '%s' line", GET_SC)
	}
	scLine = scLine[index+2:]
	tmpSlice := strings.Split(string(scLine), ", ")
	var result []int

	for _, sc := range tmpSlice {
		i, err := strconv.Atoi(sc)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, i)
	}

	return result
}

func printAndSaveTestResults(h hrm.HRM, apiTitle string) {
	now := time.Now()
	timestamp := now.Format("2006-01-02 15:04:05")
	fmt.Println("")
	fmt.Println("Date:", timestamp)
	fmt.Println(apiTitle)
	//timestamp = str(datetime.datetime.now()).replace(' ', 'T')
	resultString := ""
	header3 := "<system-out><![CDATA["

	if len(h.FailedRequestsList) > 0 {
		header := `<testsuite errors="1" failures="0" skipped="0" tests="1" timestamp="` + timestamp + `">`
		header2 := `<testcase status="failed" name="' + api_title + '">`
		header21 := `<error message="Test failed"></error>`
		resultString = header + header2 + header21 + header3
		// TODO:
		// if len(maker.warning_requests_list) > 0:
		//     mp.my_print('')
		//     mp.my_print('REQUESTS WHICH EXCEEDED WARNING TIMEOUT:')
		//     for r in maker.warning_requests_list:
		//         mp.my_print(r)
		//     mp.my_print('')
		//     if len(maker.failed_requests_list) == 0:
		//         mp.my_print('*** Test result: Warning ***')
		// if len(maker.failed_requests_list) > 0:
		tmp := "FAILED REQUESTS:\n"
		fmt.Print(tmp)
		resultString = resultString + tmp

		for _, r := range h.FailedRequestsList {
			tmp = r + "\n"
			fmt.Print(tmp)
			resultString = resultString + tmp
		}
		tmp = "\n!!! TEST FAIL !!!\n"
		fmt.Print(tmp)
		resultString = resultString + tmp
	} else {
		header := `<testsuite errors="0" failures="0" skipped="0" tests="1" timestamp="` + timestamp + `">`
		header2 := `<testcase status="passed" name="' + api_title + '">`
		resultString = header + header2 + header3
		fmt.Println("*** Test Pass ***")
		resultString = resultString + "*** Test Pass ***\n"
	}
	end := "]]></system-out></testcase></testsuite>"
	resultString = resultString + end
	filename := strings.Replace(apiTitle, " ", "_", -1) + "_test_results.xml"
	saveStringToFile(filename, resultString)
	fmt.Println("")
	fmt.Println("")
}

func saveStringToFile(filename string, str string) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(str)
	if err != nil {
		log.Fatal(err)
	}
	writer.Flush()
}

func myLog(msg string) {
	const DEBUG bool = false
	if DEBUG {
		fmt.Println("log:", msg)
	}
}
