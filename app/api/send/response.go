package send

type Option func(response Response) Response

type ResponseError struct {
	Code    string      `json:"code,omitempty"`
	Message string      `json:"message,omitempty"`
	Target  interface{} `json:"target,omitempty"`
}

type Response struct {
	error  error
	code   int
	header map[string]string
	data   interface{}
}
