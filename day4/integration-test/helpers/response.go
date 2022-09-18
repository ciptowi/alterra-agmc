package helpers

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ResponseLogin struct {
	Success     bool        `json:"success"`
	Message     string      `json:"message"`
	AccessToken string      `json:"accessToken"`
	TokenType   string      `json:"tokenType"`
	Data        interface{} `json:"data"`
}
