# Introduction
API Smoke test.
Tests GET endpoints, if resposnse status code isn't in config.txt file test fails.

# Run

`api-smoke-test-go.exe -oas OAS_file [options]`

  -oas Required argument, file name of the OpenAPI v.3 specification file.

  -auth
        Use authentication, i.e. authentication token is used to authorize requests.
        Env. variable auth_token must be set, e.g. on Windows run

`set auth_token=your_access_token`

  -help
        Show help.

  -req-param 
        Integer used in requests that contain parameters (default 13).

  -validate
        Validate file given as "oas" argument.

Result of the test will be saved to file named api_title__test_results.xml. The file has JUnit format, for example it can be used in Azure pipeline as a test result.

# TODO
-localhost Use when testing API that runs on your local machine.

-only-get Test only GET requests.

"oas" accepts url of OAS file.
