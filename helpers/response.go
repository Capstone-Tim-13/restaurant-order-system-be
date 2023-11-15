package helpers

type BaseResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type BaseSuccessResponse struct {
	Response BaseResponse `json:"response"`
	Results  interface{}  `json:"results"`
}

type BaseErrorResponse struct {
	Response BaseResponse `json:"response"`
}

func SuccessResponse(message string, data interface{}) interface{} {
	if data == nil {
		return BaseErrorResponse{
			Response: BaseResponse{
				Success: true,
				Message: message,
			},
		}
	} else {
		return BaseSuccessResponse{
			Response: BaseResponse{
				Success: true,
				Message: message,
			},
			Results: data,
		}
	}
}

func ErrorResponse(message string) interface{} {
	return BaseErrorResponse{
		Response: BaseResponse{
			Success: false,
			Message: message,
		},
	}
}
