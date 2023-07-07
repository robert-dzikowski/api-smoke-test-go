package hrm

import (
	"fmt"

	"golang.org/x/exp/slices"
)

const TIMEOUT = 5.0

// Correct HTTP status codes for GET methods
var GET_SC = []int{200, 204, 401, 403, 404}

//const POST_SC: (u16, u16, u16, u16, u16, u16) = (200, 201, 202, 204, 400, 404);

type HRM struct {
	baseApiUrl         string
	authToken          string
	FailedRequestsList []string
}

func New(baseApiURL string, authToken string) HRM {
	h := HRM{
		baseApiURL,
		authToken,
		[]string{},
	}
	return h
}

func (h *HRM) MakeGETRequests(endpoints []string) {
	var responseSc int
	for _, ep := range endpoints {
		fmt.Println("Requesting GET", ep)
		responseSc = h.sendGETRequest(h.baseApiUrl + ep)
		//fmt.Println("Status code:", responseSc)
		requestSucceeded := slices.Contains(GET_SC, responseSc)

		if !requestSucceeded {
			h.FailedRequestsList = append(
				h.FailedRequestsList,
				fmt.Sprintf("GET %s, sc: %d", ep, responseSc))
			fmt.Printf(
				"FAIL: %s request failed. Status code: %d\n", ep, responseSc)
		}
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
	}
}

func (h HRM) sendGETRequest(endPoint string) int {
	var responseSC int
	if h.authToken != "" {
		responseSC = GETProtectedResourceStatusCode(endPoint, h.authToken)
	} else {
		responseSC = GETResourceStatusCode(endPoint, nil, nil, 3)
	}
	return responseSC
}

// def _add_to_warning_list_if_exceeded_warning_timeout(self, elapsed_time, end_point):
// 	if elapsed_time > config.WARNING_TIMEOUT:
// 		self.warning_requests_list.append('GET ' + end_point)
