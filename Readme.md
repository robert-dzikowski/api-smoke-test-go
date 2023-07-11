# Introduction
Smoke test for testing APIs that use OpenAPI 3.0 specification.
Tests GET requests, if response status code isn't in config.txt file the test will fail, but all endpoints will be tested.
This prgram was tested with GitHub REST API.

# Run

`api-smoke-test-go.exe -oas OAS_file [options]`

-oas File name of the OpenAPI 3.0 specification file.
  
Options:

-auth
        Use authentication, i.e. authentication token is used to authorize requests.
        Env. variable `auth_token` must be set, e.g. on Windows run

`set auth_token=your_access_token`

  -help
        Show help.

  -req-param 
        Integer used in requests that contain parameters (default 13).

  -single-thread Use single thread for HTTP requests. By default every request is made in separate goroutine.
  
  -validate
        Validate file given as "oas" argument.

Result of the test is saved to file named api_title_test_results.xml. The file has JUnit format, so it can be used in Azure pipeline as a test result.

# TODO
Make other HTTP requests, i.e. POST, PUT, PATCH, DELETE.

oas argument accepts url of OAS file.

Add oAuth2 authorization. Currently only access token is supported.

-localhost, used when testing API that runs on your local machine.

-only-get, test only GET requests.
