package models

type HttpResponse struct {
	Status string      `json:"status_text"`
	Error  string      `json:"error_desc"`
	Data   interface{} `json:"data"`
}
