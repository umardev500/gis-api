package model

type Meta struct {
	Total int64 `json:"total"`
}

type Response struct {
	Success bool        `json:"success"`
	Status  int64       `json:"status"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
	Message string      `json:"message,omitempty"`
	Meta    *Meta       `json:"meta,omitempty"`
}
