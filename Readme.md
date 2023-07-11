# Introduction
Smoke test for testing APIs that use OpenAPI v.3 specification.
Tests GET requests, if response status code isn't in config.txt file the test will fail.

# Run

`api-smoke-test-go.exe -oas OAS_file [options]`

-oas File name of the OpenAPI v.3 specification file.
  
Options:

-auth
        Use authentication, i.e. authentication token is used to authorize requests.
        Env. variable `auth_token` must be set, e.g. on Windows run

`set auth_token=your_access_token`

  -help
        Show help.

  -req-param 
        Integer used in requests that contain parameters (default 13).

  -validate
        Validate file given as "oas" argument.

Result of the test will be saved to file named api_title_test_results.xml. The file has JUnit format, for example it can be used in Azure pipeline as a test result.

# TODO
Make other HTTP requests, i.e. POST, PUT, PATCH, DELETE.

-localhost Use when testing API that runs on your local machine.

-only-get Test only GET requests.

"oas" accepts url of OAS file.
