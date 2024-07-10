package response

type Response struct {
	Ok      bool        `json:"ok"`
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
}
