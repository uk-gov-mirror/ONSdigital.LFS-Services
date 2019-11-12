package api

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type BatchHandler struct {
}

func NewBatchHandler() *BatchHandler {
	return &BatchHandler{}
}

func intConversion(year string) int {
	yr, err := strconv.Atoi(year)
	if err != nil {
		return -1
	}
	return yr
}

func (b BatchHandler) CreateMonthlyBatchHandler(w http.ResponseWriter, r *http.Request) {
	// Variables

	vars := mux.Vars(r)
	year := vars["year"]
	month := vars["month"]
	description := r.FormValue("description")

	// Convert year to int
	yr := intConversion(year)
	if yr == -1 {
		ErrorResponse{
			Status:       Error,
			ErrorMessage: fmt.Sprintf("invalid year: %s, expected an integer", year)}.sendResponse(w, r)
		return
	}

	// Convert month to int
	mth := intConversion(month)
	if mth == -1 {
		ErrorResponse{
			Status:       Error,
			ErrorMessage: fmt.Sprintf("invalid period: %s, expected one of 1-12", month)}.sendResponse(w, r)
		return
	}

	res := b.generateMonthBatchId(mth, yr, description)
	if res != nil {
		ErrorResponse{Status: Error, ErrorMessage: res.Error()}.sendResponse(w, r)
		return
	}

	OkayResponse{OK}.sendResponse(w, r)

}

func (b BatchHandler) CreateQuarterlyBatchHandler(w http.ResponseWriter, r *http.Request) {
	// Variables
	vars := mux.Vars(r)
	year := vars["year"]
	quarter := vars["quarter"]
	description := r.FormValue("description")

	// Convert year to int
	yr := intConversion(year)
	if yr == -1 {
		ErrorResponse{
			Status:       Error,
			ErrorMessage: fmt.Sprintf("invalid year: %s, expected an integer", year)}.sendResponse(w, r)
		return
	}

	// Strip and convert period to int
	q, err := strconv.Atoi(quarter[1:])
	if err != nil || len(quarter) != 2 {
		ErrorResponse{
			Status:       Error,
			ErrorMessage: fmt.Sprintf("invalid period: %s, expected one of Q1-Q4", quarter)}.sendResponse(w, r)
		return
	}

	res := b.generateQuarterBatchId(p, yr, description)
	if res != nil {
		ErrorResponse{Status: Error, ErrorMessage: res.Error()}.sendResponse(w, r)
		return
	}

	OkayResponse{OK}.sendResponse(w, r)
}

func (b BatchHandler) CreateAnnualBatchHandler(w http.ResponseWriter, r *http.Request) {
	// Variables
	vars := mux.Vars(r)
	year := vars["year"]
	description := r.FormValue("description")

	// Convert year to int
	yr := intConversion(year)
	if yr < -1 || yr == 0 {
		ErrorResponse{
			Status:       Error,
			ErrorMessage: fmt.Sprintf("invalid year: %s, expected an integer", year)}.sendResponse(w, r)
		return
	}

	res := b.generateYearBatchId(yr, description)
	if res != nil {
		ErrorResponse{Status: Error, ErrorMessage: res.Error()}.sendResponse(w, r)
		return
	}

	OkayResponse{OK}.sendResponse(w, r)
}
