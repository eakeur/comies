package send

type Option func(response Response) Response

type ResponseError struct {
	Code    string      `json:"code,omitempty"`
	Message string      `json:"message,omitempty"`
	Target  interface{} `json:"target,omitempty"`
}

type Response struct {
	Error  error
	Code   int
	Header map[string]string
	Data   interface{}
}
