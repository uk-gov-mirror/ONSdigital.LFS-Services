package api

import (
	"encoding/json"
	"github.com/rs/zerolog/log"
	"net/http"
	"time"
)

type Response interface {
	sendResponse(w http.ResponseWriter, r *http.Request)
}

type DataResponse interface {
	sendDataResponse(w http.ResponseWriter, r *http.Request, d interface{})
}

type UnknownFileType struct {
	Status       string `json:"status"`
	ErrorMessage string `json:"errorMessage"`
}

type ErrorResponse struct {
	Status       string `json:"status"`
	ErrorMessage string `json:"errorMessage"`
}

type BadDataResponse struct {
	Status       string `json:"status"`
	ErrorMessage string `json:"errorMessage"`
	Result       interface{}
}

type OkayResponse struct {
	Status string `json:"status"`
}

type SendDataResponse struct{}

type NoRecordsFoundStatus struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type InProgressResponse struct {
	Status  string `json:"status"`
	When    string `json:"time"`
	Message string `json:"message"`
}

func (re NoRecordsFoundStatus) sendResponse(w http.ResponseWriter, r *http.Request) {
	re.Status = OK
	re.Message = "no data found"
	w.WriteHeader(http.StatusOK)
	sendResponse(w, r, re)
}

func (re SendDataResponse) sendResponse(w http.ResponseWriter, r *http.Request, d interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(d); err != nil {
		log.Error().
			Str("client", r.RemoteAddr).
			Str("uri", r.RequestURI).
			Msg("json.NewEncoder() failed in sendDataResponse")
	}
}

func (response InProgressResponse) sendResponse(w http.ResponseWriter, r *http.Request) {
	response.Status = OK
	response.When = time.Now().String()
	response.Message = "request submitted"
	w.WriteHeader(http.StatusAccepted)
	sendResponse(w, r, response)
}

func (response OkayResponse) sendResponse(w http.ResponseWriter, r *http.Request) {
	response.Status = OK
	w.WriteHeader(http.StatusOK)
	sendResponse(w, r, response)
}

type badDataesult struct {
	Status       string `json:"status"`
	ErrorMessage string `json:"errorMessage"`
	Result       interface{}
}

func (response BadDataResponse) sendResponse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	s := badDataesult{
		Status:       response.Status,
		ErrorMessage: response.ErrorMessage,
		Result:       response.Result,
	}
	if err := json.NewEncoder(w).Encode(s); err != nil {
		log.Error().
			Str("client", r.RemoteAddr).
			Str("uri", r.RequestURI).
			Msg("json.NewEncoder() failed in sendBadDataResponse")
	}
}

func (response ErrorResponse) sendResponse(w http.ResponseWriter, r *http.Request) {
	response.Status = Error
	w.WriteHeader(http.StatusBadRequest)
	sendResponse(w, r, response)
}

func (response UnknownFileType) sendResponse(w http.ResponseWriter, r *http.Request) {
	response.Status = Error
	w.WriteHeader(http.StatusBadRequest)
	sendResponse(w, r, response)
}

func sendResponse(w http.ResponseWriter, r *http.Request, response Response) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Error().
			Str("client", r.RemoteAddr).
			Str("uri", r.RequestURI).
			Msg("json.NewEncoder() failed in FileUploadHandler")
	}
}
