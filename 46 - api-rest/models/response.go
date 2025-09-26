package models

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Status      int    `json:"status"`
	Data        any    `json:"data"`
	Message     string `json:"message"`
	contentType string
	respWrite   http.ResponseWriter
}

func CreateDefaultResponse(rw http.ResponseWriter) Response {
	return Response{
		Status:      http.StatusOK,
		respWrite:   rw,
		contentType: "application/json",
	}
}

func SendData(rw http.ResponseWriter, data any) {
	response := CreateDefaultResponse(rw)
	response.Data = data
	response.send()
}

func (resp *Response) NotFound() {
	resp.Status = http.StatusNotFound
	resp.Message = "Resouce not found"
}

func SendNotFound(rw http.ResponseWriter) {
	response := CreateDefaultResponse(rw)
	response.NotFound()
	response.send()
}

func (resp *Response) UnprocessableEntity() {
	resp.Status = http.StatusUnprocessableEntity
	resp.Message = "UnprocessableEntity"
}

func SendUnprocessableEntity(rw http.ResponseWriter) {
	response := CreateDefaultResponse(rw)
	response.UnprocessableEntity()
	response.send()
}

func (resp *Response) send() {
	resp.respWrite.Header().Set("Content-Type", resp.contentType)
	resp.respWrite.WriteHeader(resp.Status)

	byteResult, _ := json.Marshal(&resp)
	fmt.Fprintln(resp.respWrite, string(byteResult))
}
