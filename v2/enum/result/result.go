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

func NewWithCodeAndError(code Code, err error, data interface{}) *Result {
	return &Result{
		C:       code,
		Message: code.String() + ": " + err.Error(),
		Data:    data,
	}
}

func NewWithCodeAndData(code Code, data interface{}) *Result {
	return &Result{
		C:       code,
		Message: code.String(),
		Data:    data,
	}
}

func NewWithCode(code Code) *Result {
	return &Result{
		C:       code,
		Message: code.String(),
	}
}