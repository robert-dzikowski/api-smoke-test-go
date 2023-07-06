package hrm

import (
	"log"
	"net/http"
	"time"
)

const TIMEOUT = 10.0

func CheckError(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

// Send GET request, return status code of reply
func GETResourceStatusCode(
	endpoint string, params map[string]string, headers map[string]string) int {
	resp, err := getResource(endpoint, params, headers)
	CheckError(err)
	return resp.StatusCode
}

func getResource(
	endpoint string, params map[string]string,
	headers map[string]string) (*http.Response, error) {
	tries := 0

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

	for {
		tries += 1
		resp, err := client.Do(req)
		if err == nil || tries >= 3 {
			return resp, err
		}
	}
}

// func create_408_response(error_msg string) (*http.Response, error) {
// 	resp := &http.Response{
// 		StatusCode: http.StatusRequestTimeout,
// 		Body:       http.NoBody,
// 	}
// 	resp.Header = make(http.Header)
// 	resp.Header.Set("Content-Type", "text/plain")
// 	resp.Write([]byte(error_msg))
// 	return resp, fmt.Errorf(error_msg)
// }
