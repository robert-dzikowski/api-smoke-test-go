package hrm

import (
	"fmt"
	"sort"

	"golang.org/x/exp/slices"
)

type HRM struct {
	baseApiUrl         string
	authToken          string
	Timeout            float64
	GetSC              []int
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

func (h *HRM) MakeGETRequests(endpoints []string) {
	c := make(chan string)

	var responseSc int
	for _, ep := range endpoints {
		endPoint := ep
		go func() {
			fmt.Println("Requesting GET", endPoint)
			responseSc = h.sendGETRequest(h.baseApiUrl + endPoint)
			requestSucceeded := slices.Contains(h.GetSC, responseSc)

			if requestSucceeded {
				c <- ""
			} else {
				// ,
				c <- fmt.Sprintf("GET %s, sc: %d", endPoint, responseSc)
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
	sort.Strings(h.FailedRequestsList)
}

func (h HRM) sendGETRequest(endPoint string) int {
	var responseSC int
	if h.authToken != "" {
		responseSC = GETProtectedResourceStatusCode(endPoint, h.authToken, h.Timeout)
	} else {
		responseSC = GETResourceStatusCode(endPoint, h.Timeout, 3)
	}
	return responseSC
}

// def _add_to_warning_list_if_exceeded_warning_timeout(self, elapsed_time, end_point):
// 	if elapsed_time > config.WARNING_TIMEOUT:
// 		self.warning_requests_list.append('GET ' + end_point)
