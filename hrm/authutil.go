package hrm

// import (
// )

import (
	"fmt"
	"net/http"
	"net/url"
	"time"
	//"context"
	//"golang.org/x/oauth2"
)

func GETProtectedResourceStatusCode(endpoint string, token string) int {
	resp, err := getProtectedResource(endpoint, token)
	CheckError(err)
	return resp.StatusCode
}

func getProtectedResource(endpoint string, token string) (*http.Response, error) {
	client := http.Client{
		Timeout: time.Duration(TIMEOUT * float64(time.Second)),
	}
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Token %s", token))
	resp, err := client.Do(req)
	if urlErr, ok := err.(*url.Error); ok && urlErr.Timeout() {
		resp = get408Response()
		err = nil
	}

	if err != nil {
		return nil, err
	}
	// TODO:
	// elapsed_time := response.elapsed.total_seconds()
	// print(' Duration: ' + str(elapsed_time))

	return resp, nil
}

// func getProtectedResource(
// 	endpoint string, clientId string, token *oauth2.Token,
// 	getTimeout float64, headers map[string]string) (*http.Response, error) {
// 	ctx := context.Background()
// 	client := oauth2.NewClient(ctx, oauth2.StaticTokenSource(token))
// 	client.Timeout = time.Duration(getTimeout * float64(time.Second))
// 	req, err := http.NewRequest("GET", endpoint, nil)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if headers != nil {
// 		for key, value := range headers {
// 			req.Header.Set(key, value)
// 		}
// 	}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return resp, nil
// }
