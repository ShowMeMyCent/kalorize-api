package utils

type Response struct {
	StatusCode int         `json:"statusCode"`
	Messages   string      `json:"messages"`
	Data       interface{} `json:"data"`
}
