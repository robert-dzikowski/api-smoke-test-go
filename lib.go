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
	ctx := context.Background()
	loader := &openapi3.Loader{Context: ctx}

	// 2.1 Read and parse OAS file
	// let data: OpenAPI;
	var doc *openapi3.T
	var err error

	if strings.HasPrefix(*args.oasFile, "http") {
		parsedURL, e := url.Parse(*args.oasFile)
		check(e)
		doc, err = loader.LoadFromURI(parsedURL)
	} else {
		doc, err = loader.LoadFromFile(*args.oasFile)
	}

	check(err)

	// Validate document
	if *args.validate {
		err = doc.Validate(ctx)
		check(err)
	}

	fmt.Println("Title:", doc.Info.Title)

	// let base_api_url = (&data.servers[0].url).to_string();
	// println!("Base URL: {:?}", base_api_url);
	// println!("");
	// println!("Testing {}", data.info.title); //['info']['title'])
	// println!("");

	// // 3. Create list of GET endpoints.
	// let endpoints_list = return_list_of_parameterless_get_methods(data);
	// utilities::my_log(format!("GET endpoints: {:?}", endpoints_list));

	// // endpoints_with_params = return_list_of_get_methods_with_parameters(
	// //     paths_dict)

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
