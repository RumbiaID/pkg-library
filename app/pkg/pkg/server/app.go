package server

const (
	DefaultResponseType = iota // For API
	StreamResponseType         // For Excel
)

type ResponseInterface struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}

type Response struct {
	Status       int         `json:"status"`
	Message      string      `json:"message"`
	Data         interface{} `json:"data"`
	RequestID    string      `json:"request_id"`
	StackTrace   string      `json:"-"`
	ResponseType int         `json:"-"`
}

type ResponseSignature struct {
	Signature string `json:"signature"`
}

func (r Response) GetStatus() int {
	return r.Status

}
