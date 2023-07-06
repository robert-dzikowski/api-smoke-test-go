package hrm

import "fmt"

const TIMEOUT = 10.0

// Correct HTTP status codes for GET methods
var GET_SC = [3]int{200, 204, 400}

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
	//} else {
	// response = utils.get_resource(
	//     self._api_url + end_point, headers=self.headers)
	//}
}
