package hrm

import "fmt"

const TIMEOUT = 10.0

// Correct HTTP status codes for GET methods
var GET_CORRECT = [5]int{200, 204, 401, 403, 404}

//const POST_SC: (u16, u16, u16, u16, u16, u16) = (200, 201, 202, 204, 400, 404);

type HRM struct {
	baseApiUrl string
	authToken  string
}

func New(baseApiURL string, authToken string) HRM {
	h := HRM{
		baseApiURL,
		authToken,
	}
	return h
}

func (h HRM) MakeGETRequests(endpoints []string) {
	var responseSC int
	for _, ep := range endpoints {
		fmt.Println("Requesting GET", ep)
		responseSC = h.sendGETRequest(h.baseApiUrl + ep)
		fmt.Println("Status code:", responseSC)
		// request_succeeded = (status_code in correct_statuses)
		// if not request_succeeded:
		// 	self.failed_requests_list.append(
		// 		http_method.value + ' ' + end_point + ', sc: ' + str(status_code))
		// 	print('FAIL: ' + end_point +
		// 		  ' request failed. Status code: ' + str(status_code))
		// 	print('')
		// else:
		// 	if http_method == HttpMethods.GET:
		// 		self._add_to_warning_list_if_exceeded_warning_timeout(
		// 			elapsed_time, end_point)
		// 	else:
		// 		self._add_to_warning_list_if_exceeded_warning_timeout_post(
		// 			elapsed_time, end_point, http_method)
	}
}

func (h HRM) sendGETRequest(endPoint string) int {
	var responseSC int
	if h.authToken != "" {
		responseSC = GETProtectedResourceStatusCode(endPoint, h.authToken)
	} else {
		responseSC = GETResourceStatusCode(endPoint, nil, nil)
	}
	return responseSC
}

// def _add_to_warning_list_if_exceeded_warning_timeout(self, elapsed_time, end_point):
// 	if elapsed_time > config.WARNING_TIMEOUT:
// 		self.warning_requests_list.append('GET ' + end_point)

// def _add_to_warning_list_if_exceeded_warning_timeout_post(
// 		self, elapsed_time, end_point, http_method):
// 	if elapsed_time > config.WARNING_TIMEOUT_POST:
// 		self.warning_requests_list.append(
// 		http_method.value + ' ' + end_point)
