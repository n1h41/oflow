package model

// INFO: Requests

// INFO: Responses

type GlobalErrorHandlerResp struct {
	Status  bool        `json:"status"`
	Message interface{} `json:"message"`
}
