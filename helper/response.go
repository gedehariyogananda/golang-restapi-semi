package helper

type ApiResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Status  bool        `json:"status"`
	Code    int         `json:"code"`
}
