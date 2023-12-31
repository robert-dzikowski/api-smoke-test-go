package hrm

import (
	"fmt"
	"net/http"
	"net/url"
	"time"
)

func GETProtectedResourceStatusCode(
	endpoint string, token string, timeout float64) int {
	resp, err := GETProtectedResource(endpoint, token, timeout)
	CheckError(err)
	return resp.StatusCode
}

func GETProtectedResource(
	endpoint string, token string, timeout float64) (*http.Response, error) {
	client := http.Client{
		Timeout: time.Duration(timeout * float64(time.Second)),
	}
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Token %s", token))
	resp, err := client.Do(req)

	urlErr, ok := err.(*url.Error)
	if ok && urlErr.Timeout() {
		resp = get408Response()
		err = nil
	}

	// if strings.Contains(endpoint, "/albums") {
	// 	fmt.Println("Endpoint:", endpoint)

	// TODO:
	// elapsed_time := response.elapsed.total_seconds()
	// print(' Duration: ' + str(elapsed_time))

	return resp, err
}
