package todoModels

type (
	ResponseMessage struct {
		Message string `json:"message"`
		Status  int    `json:"status"`
	}
	Response struct {
		Data   any `json:"data"`
		Status int `json:"status"`
	}
	ErrorRespnse struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}
)
