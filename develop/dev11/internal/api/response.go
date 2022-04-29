package api

type Result struct {
	Res        string `json:"result"`
	StatusCode int    `json:"statusCode"`
}
type Err struct {
	ErrResponse string `json:"error"`
	StatusCode  int    `json:"statusCode"`
}

func NewResult(message string, statusCode int) *Result {
	return &Result{
		Res:        message,
		StatusCode: statusCode,
	}
}
func NewErr(message string, statusCode int) *Err {
	return &Err{
		ErrResponse: message,
		StatusCode:  statusCode,
	}
}
