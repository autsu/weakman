package result

type Result struct {
	C       Code        `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func New(code Code, message string, data interface{}) *Result {
	return &Result{
		C:       code,
		Message: message,
		Data:    data,
	}
}

