package response

type WebResponse struct {
	Code   int         `json:"statusCode"`
	Status string      `json:"status"`
	Data   interface{} `json:"data,omitempty"`
}
