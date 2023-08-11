package web

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

const (
	// Header Keys
	ContentType        = "Content-Type"
	ContentDisposition = "Content-Disposition"
	ContentLength      = "Content-Length"
	ListCount          = "List-Count"

	// Header Values
	ApplicationJSON    = "application/json"
	ApplicationCSV     = "application/csv"
	ApplicationPDF     = "application/pdf"
	TextPlain          = "text/plain"
	TextCalendar       = "text/calendar"
	HTMLUTF8           = "text/html; charset=utf-8"
	AttachmentFilename = "attachment; filename=%s.%s"

	// File Extensions
	CSV = "csv"
	PDF = "pdf"
	TXT = "txt"
)

type APIError struct {
	ErrorMessage error `json:"errorMessage"`
}

// Response Helpers

func attachment(filename, extension string) string {
	return fmt.Sprintf(AttachmentFilename, filename, extension)
}

func RespondString(w http.ResponseWriter, status int, s string) {
	w.WriteHeader(status)
	w.Write([]byte(s))
}

func RespondPlainText(w http.ResponseWriter, s string) {
	w.Header().Set(ContentType, TextPlain)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(s))
}

func RespondHTML(w http.ResponseWriter, html []byte) {
	w.Header().Set(ContentType, HTMLUTF8)
	w.Write(html)
}

func RespondCalendar(w http.ResponseWriter, s string) {
	w.Header().Set(ContentType, TextCalendar)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(s))
}

func RespondXML(w http.ResponseWriter, xml []byte) {
	xml = append([]byte(`<?xml version="1.0" encoding="utf-8"?>`), xml...)
	w.Header().Set("Content-Type", "application/xhtml+xml")
	w.Write(xml)
}

func RespondOK(w http.ResponseWriter, content interface{}) {
	w.Header().Set(ContentType, ApplicationJSON)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(content)
}

func RespondBlank(w http.ResponseWriter) {
	RespondOK(w, nil)
}

func RespondNothing(w http.ResponseWriter, status int) {
	w.WriteHeader(status)
}

func RespondError(w http.ResponseWriter, err error, status int) {
	w.Header().Set(ContentType, ApplicationJSON)
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(APIError{err})
}

func RespondBadRequest(w http.ResponseWriter, err error) {
	RespondError(w, err, http.StatusBadRequest)
}

func RespondUnauthorized(w http.ResponseWriter, err error) {
	RespondError(w, err, http.StatusUnauthorized)
}

func RespondNotFound(w http.ResponseWriter, err error) {
	RespondError(w, err, http.StatusNotFound)
}

func RespondOKWithCount(w http.ResponseWriter, content interface{}, count int64) {
	w.Header().Set(ListCount, fmt.Sprintf("%v", count))
	RespondOK(w, content)
}

func RespondCSVEntityList(w http.ResponseWriter, records [][]string, filename string) {
	w.Header().Set(ContentDisposition, attachment(filename, CSV))
	w.Header().Set(ContentType, ApplicationCSV)

	csvWriter := csv.NewWriter(w)
	csvWriter.WriteAll(records) // calls Flush internally

	if err := csvWriter.Error(); err != nil {
		RespondError(w, err, http.StatusInternalServerError)
	}
}

func RespondPDFFile(w http.ResponseWriter, filename string, pdfOutput []byte) {
	RespondFile(w, filename, PDF, ApplicationPDF, pdfOutput)
}

func RespondFile(w http.ResponseWriter, filename, extension, contentType string, contents []byte) {
	w.Header().Set(ContentDisposition, attachment(filename, extension))
	w.Header().Set(ContentType, contentType)
	w.Header().Set(ContentLength, strconv.Itoa(len(contents)))
	w.Write(contents)
}

type BasicAuthCredentials struct {
	Username, Password, Realm string
}

func RespondBasicAuthUnauthorized(w http.ResponseWriter, credentials *BasicAuthCredentials) {
	w.Header().Set("WWW-Authenticate", `Basic realm="`+credentials.Realm+`"`)
	w.WriteHeader(401)
	w.Write([]byte("Unauthorised.\n"))
}
