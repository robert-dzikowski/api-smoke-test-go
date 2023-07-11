# Introduction
Smoke test for testing APIs that use OpenAPI v.3 specification.
Tests GET requests, if response status code isn't in config.txt file the test will fail.
Tested with GitHub REST API.

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

  -single-thread Use single thread for HTTP requests. By default every request is made in separate goroutine.
  
  -validate
        Validate file given as "oas" argument.

Result of the test will is saved to file named api_title_test_results.xml. The file has JUnit format, so it can be used in Azure pipeline as a test result.

# TODO
Make other HTTP requests, i.e. POST, PUT, PATCH, DELETE.

oas argument accepts url of OAS file.

Add oAuth2 authorization. Curremtly only access token is supported, i.e. you have to get your access token manually if tested API uses oAuth2.

-localhost, used when testing API that runs on your local machine.

-only-get, test only GET requests.
