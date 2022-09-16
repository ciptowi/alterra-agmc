package lib

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message" form:"message"`
	Data    interface{} `json:"data"`
}
