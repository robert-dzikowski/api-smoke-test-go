package hrm

import (
	"fmt"
	"net/http"
	"net/url"
	"time"
)

func GETProtectedResourceStatusCode(
	endpoint string, token string, timeout float64) int {
	resp, err := getProtectedResource(endpoint, token, timeout)
	CheckError(err)
	return resp.StatusCode
}

func getProtectedResource(
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
	if urlErr, ok := err.(*url.Error); ok && urlErr.Timeout() {
		resp = get408Response()
		err = nil
	}

	defer resp.Body.Close()

	// if strings.Contains(endpoint, "/albums") {
	// 	fmt.Println("Endpoint:", endpoint)

	// 	body, _ := ioutil.ReadAll(resp.Body)
	// 	fmt.Println("Response:", string(body))
	// 	fmt.Println("")
	// }

	// TODO:
	// elapsed_time := response.elapsed.total_seconds()
	// print(' Duration: ' + str(elapsed_time))

	return resp, err
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
