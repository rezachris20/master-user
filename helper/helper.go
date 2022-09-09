package helper

type ResponseDetail struct {
	Message   string      `json:"message"`
	Data      interface{} `json:"data,omitempty"`
	IsSuccess interface{} `json:"is_success,omitempty"`
}

func ResponseDetailOutput(message string, data interface{}) ResponseDetail {
	return ResponseDetail{
		Message: message,
		Data:    data,
	}
}
