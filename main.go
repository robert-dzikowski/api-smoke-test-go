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

// func run(args: Args) {
//     // 2. Parse OAS spec from file or internet
//     let config = Config::build(args);

//     // 2.1 Read and parse OAS file
//     let data: OpenAPI;

//     if config.spec_file.starts_with("http") {
//         // content = utils.get_resource_content_string(spec_file)
//         // spec = yaml.safe_load(content)
//         data = parse_spec_file(String::from("dummy")); // added this line to satisfy compiler
//     } else {
//         let contents = fs::read_to_string(&config.spec_file).expect(
//             format!("Error opening file {}", config.spec_file).as_str(),
//         );
//         data = parse_spec_file(contents);
//     }

//     let base_api_url = (&data.servers[0].url).to_string();
//     println!("Base URL: {:?}", base_api_url);
//     println!("");
//     println!("Testing {}", data.info.title); //['info']['title'])
//     println!("");

//     // 3. Create list of GET endpoints.
//     let endpoints_list = return_list_of_parameterless_get_methods(data);
//     utilities::my_log(format!("GET endpoints: {:?}", endpoints_list));

//     // endpoints_with_params = return_list_of_get_methods_with_parameters(
//     //     paths_dict)

//     // 4. If endpoints_list list is empty exit.
//     if endpoints_list.len() == 0 {
//         println!(
//             "Test failed: spec file {} doesn't contain any GET methods.",
//             config.spec_file
//         );
//         process::exit(1);
//     }

//     // 5. Get auth token
//     // token = None
//     // if authorization_is_necessary():
//     //     token = get_auth_token()

//     // 6. Test parameterless GET endpoints
//     let hrm_conf = HRM::build(base_api_url);
//     println!("Testing GET methods");

//     hrm_conf.make_get_requests(endpoints_list);

//     // 7. Test GET endpoints that contain parameters
//     // req_param = get_request_param_arg()
//     //
//     // if len(endpoints_with_params) > 0:
//     //     call_get_methods_with_parameters(
//     //         endpoints_with_params, maker, req_param)

//     // 8. Print test results.
//     // print_test_results(maker, spec['info']['title'])

//     // 9. Exit with error if any test failed.
//     // Exit with error code is needed by Azure to show test as failed
//     // if len(maker.failed_requests_list) > 0:
//     //     sys.exit(1)

//     // 10. Exit with error if any test returned warning.
//     // if config.WARNING_FAIL and len(maker.warning_requests_list) > 0:
//     //     sys.exit(1)
// }
