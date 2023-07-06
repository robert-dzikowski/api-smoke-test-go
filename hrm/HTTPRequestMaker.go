package hrm

import "fmt"

// Correct HTTP status codes for GET verbs
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
	for _, ep := range endpoints {
		fmt.Println("Requesting GET", ep)
		response := sendGETRequest(h.baseApiUrl + ep)
		fmt.Println("Status code:", response)
	}
}

func sendGETRequest(endPoint string) int {
	// fmt.Println(endPoint) // http://petstore.swagger.io/v2/pets
	replySC := GETResourceStatusCode(endPoint, nil, nil)
	return replySC
	//} else {
	// response = utils.get_resource(
	//     self._api_url + end_point, headers=self.headers)
	//}
}
