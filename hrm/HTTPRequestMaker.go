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

func (HRM) MakeGETRequests(endpoints []string) {
	for _, ep := range endpoints {
		fmt.Println("Requesting GET", ep)
		// 	let response = self.send_get_request(format!(
		// 		"{}{}",
		// 		self.base_api_url, end_point
		// 	));
		// 	println!("Status code: {}", response);
	}
}

//     fn send_get_request(&self, end_point: String) -> u16 {
//         // let response = Result<Response, Error>();
//         //if self.auth_token.is_none() {
//         let reply_sc = get_resource_status_code(&end_point, None);
//         reply_sc
//         //} else {
//         // response = utils.get_resource(
//         //     self._api_url + end_point, headers=self.headers)
//         //}
//     }
// }
