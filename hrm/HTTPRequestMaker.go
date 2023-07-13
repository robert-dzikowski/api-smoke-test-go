package hrm

import (
	"fmt"
	"io"
	"net/http"
	"sort"

	"golang.org/x/exp/slices"
)

type HRM struct {
	baseApiUrl         string
	authToken          string
	Timeout            float64
	ScGet              []int
	FailedRequestsList []string
}

func New(
	baseApiURL string, authToken string,
	timeout float64, getStatusCodes []int) HRM {
	h := HRM{
		baseApiURL,
		authToken,
		timeout,
		getStatusCodes,
		[]string{},
	}
	return h
}

func (h *HRM) MakeGETRequests(endpoints []string, singleThread bool) {
	if singleThread {
		h.makeGETRequestsST(endpoints)
	} else {
		h.makeGETRequests(endpoints)
	}
	sort.Strings(h.FailedRequestsList)
}

func (h *HRM) makeGETRequests(endpoints []string) {
	c := make(chan string)

	for _, ep := range endpoints {
		endPoint := ep
		go func() {
			fmt.Println("Requesting GET", endPoint)
			response := h.sendGETRequest(h.baseApiUrl + endPoint)

			defer response.Body.Close()

			responseSc := response.StatusCode
			requestSucceeded := slices.Contains(h.ScGet, responseSc)

			if requestSucceeded {
				c <- ""
			} else {
				fr := fmt.Sprintf("GET %s, sc: %d\n", endPoint, responseSc)
				fr = fr + fmt.Sprintf("Response: %s", getResponseBody(response))
				c <- fr
				fmt.Printf(
					"FAIL: %s request failed. Status code: %d\n", endPoint, responseSc)
			}
		}()
		// TODO:
		// else {
		// 	if http_method == HttpMethods.GET {
		// 		self._add_to_warning_list_if_exceeded_warning_timeout(
		// 			elapsed_time, end_point)
		// 	} else {
		// 		self._add_to_warning_list_if_exceeded_warning_timeout_post(
		// 			elapsed_time, end_point, http_method)
		// 	}
		// }
	} //for

	lenEP := len(endpoints)
	//fmt.Println("len(endpoints):", lenEP)

	for i := 0; i < lenEP; i++ {
		result := <-c
		if result != "" {
			h.FailedRequestsList = append(h.FailedRequestsList, result)
		}
	}
}

// Single thread requests
func (h *HRM) makeGETRequestsST(endpoints []string) {
	var response *http.Response
	var responseSc int
	var fr string

	for _, ep := range endpoints {
		fmt.Println("Requesting GET", ep)
		response = h.sendGETRequest(h.baseApiUrl + ep)

		defer response.Body.Close()

		responseSc = response.StatusCode
		requestSucceeded := slices.Contains(h.ScGet, responseSc)

		if !requestSucceeded {
			fr = fmt.Sprintf("GET %s, sc: %d\n", ep, responseSc)
			fr = fr + fmt.Sprintf("Response: %s", getResponseBody(response))
			h.FailedRequestsList = append(h.FailedRequestsList, fr)
			fmt.Printf(
				"FAIL: %s request failed. Status code: %d\n", ep, responseSc)
		}
	}
}

func (h HRM) sendGETRequest(endPoint string) *http.Response {
	var response *http.Response
	var err error

	if h.authToken != "" {
		response, err = GETProtectedResource(endPoint, h.authToken, h.Timeout)
	} else {
		response, err = GETResource(endPoint, h.Timeout, 3)
	}
	CheckError(err)
	return response
}

func getResponseBody(resp *http.Response) string {
	//body, _ := ioutil.ReadAll(resp.Body)
	body, _ := io.ReadAll(resp.Body)
	return string(body)
}

// def _add_to_warning_list_if_exceeded_warning_timeout(self, elapsed_time, end_point):
// 	if elapsed_time > config.WARNING_TIMEOUT:
// 		self.warning_requests_list.append('GET ' + end_point)
