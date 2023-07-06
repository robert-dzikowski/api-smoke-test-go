package hrm

import (
	"net/http"
	"time"
)

const TIMEOUT = 10.0

// Send GET request, return status code of reply
func GETResourceStatusCode(
	endpoint string, params map[string]string) int {
	// resp, err := get_resource(endpoint, params, nil)
	// return resp.status()
	return 0
}

func get_resource(
	endpoint string, params map[string]string,
	headers map[string]string) (*http.Response, error) {
	tries := 0
	for {
		tries += 1
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
