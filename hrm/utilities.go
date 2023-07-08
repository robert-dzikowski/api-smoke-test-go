package hrm

import (
	"log"
	"net/http"
	"net/url"
	"time"
)

func CheckError(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

// Send GET request, return status code of reply
func GETResourceStatusCode(
	endpoint string, params map[string]string,
	headers map[string]string, maxTries int) int {
	resp, err := getResource(endpoint, params, headers, maxTries)
	CheckError(err)
	return resp.StatusCode
}

// GET request without verification.
// maxTries > 1 is used when getting OAS file from server
func getResource(
	endpoint string,
	params map[string]string,
	headers map[string]string,
	maxTries int) (*http.Response, error) {
	tries := 1
	client := http.Client{
		Timeout: time.Duration(TIMEOUT * float64(time.Second)),
	}

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	if headers != nil {
		for key, value := range headers {
			req.Header.Set(key, value)
		}
	}

	if params != nil {
		q := req.URL.Query()
		for key, value := range params {
			q.Add(key, value)
		}
		req.URL.RawQuery = q.Encode()
	}

	// for {
	// 	tries += 1
	// 	resp, err := client.Do(req)
	// 	if err != nil {
	// 		if urlErr, ok := err.(*url.Error); ok && urlErr.Timeout() {
	// 			//fmt.Println("Timeout error")
	// 			if tries >= maxTries {
	// 				return get408Response(), nil
	// 			}
	// 		}
	// 	}
	// 	if err == nil || tries >= maxTries {
	// 		return resp, err
	// 	}
	// }
	var resp *http.Response

	for tries <= maxTries {
		resp, err = client.Do(req)
		if err == nil {
			return resp, err
		}
		tries++
	}

	if isTimeoutError(err) {
		return get408Response(), nil
	} else {
		return resp, err
	}
}

func isTimeoutError(err error) bool {
	urlErr, ok := err.(*url.Error)
	return ok && urlErr.Timeout()
}

func get408Response() *http.Response {
	result := http.Response{
		StatusCode: 408,
	}
	return &result
}
