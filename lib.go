package main

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
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
		check(e)
		doc, err = loader.LoadFromURI(parsedURL)
	} else {
		doc, err = loader.LoadFromFile(*args.oasFile) //("specs/petstore.json") //(*args.oasFile)
	}

	check(err)

	// Validate OAS document
	if *args.validate {
		err = doc.Validate(ctx)
		check(err)
	}

	baseApiUrl := doc.Servers[0].URL
	fmt.Println("Base URL:", baseApiUrl)
	fmt.Println("")
	fmt.Println("Testing", doc.Info.Title)
	fmt.Println("")

	// 3. Create list of GET endpoints
	endpointsList := getListOfParameterlessGETMethods(doc)
	myLog(fmt.Sprint("Parmeterless GET endpoints: ", endpointsList))

	// endpoints_with_params = return_list_of_get_methods_with_parameters(
	//     paths_dict)

	// // 4. If endpoints_list list is empty exit.
	// if endpoints_list.len() == 0 {
	//     println!(
	//         "Test failed: spec file {} doesn't contain any GET methods.",
	//         config.spec_file
	//     );
	//     process::exit(1);
	// }

	// // 5. Get auth token
	// // token = None
	// // if authorization_is_necessary():
	// //     token = get_auth_token()

	// // 6. Test parameterless GET endpoints
	// let hrm_conf = HRM::build(base_api_url);
	// println!("Testing GET methods");

	// hrm_conf.make_get_requests(endpoints_list);

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

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func getListOfParameterlessGETMethods(oasDoc *openapi3.T) []string {
	result := []string{}

	for path, pathItem := range oasDoc.Paths {
		// if endpoint_has_get_method(path_item) {
		//     tmp = path.to_string();
		//     if !tmp.contains('{') {
		//         result.push(tmp);
		//     }
		// }
		//fmt.Println(path)
		for method := range pathItem.Operations() {
			if method == "GET" && !strings.Contains(path, "{") {
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
