package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const (
	applicationJson = "application/json"

	UnsupportedMediaType = ErrorCode(1)
	MalformedInput       = ErrorCode(2)
)

type ErrorCode int

type Error struct {
	Message string
	Code    ErrorCode
}

func (e Error) String() string {
	return fmt.Sprintf("Payload parse error: %s", e.Message)
}

func (e Error) Error() string {
	return e.String()
}

func WriteData(w http.ResponseWriter, httpCode int, d interface{}) {
	WriteJSON(w, httpCode, d)
}

func WriteError(w http.ResponseWriter, httpCode int, d interface{}) {
	WriteJSON(w, httpCode, d)
}

func WriteJSON(w http.ResponseWriter, httpCode int, d interface{}) {
	j, err := json.Marshal(d)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpCode)
	_, _ = w.Write(j)
}

func Payload(r *http.Request, value interface{}) error {
	contentType := strings.ToLower(getContentType(r.Header))
	if contentType == "" {
		contentType = applicationJson
	}
	if strings.Contains(contentType, applicationJson) {
		return parseJson(r.Body, value)
	}

	return Error{Code: UnsupportedMediaType, Message: fmt.Sprintf("Unsupported Content-Type '%s'", contentType)}
}

func parseJson(r io.ReadCloser, parsed interface{}) error {
	err := json.NewDecoder(r).Decode(parsed)
	defer r.Close()

	if err != nil {
		return Error{Message: err.Error(), Code: MalformedInput}
	}
	return nil
}

func getContentType(headers http.Header) string {
	for key, value := range headers {
		if strings.Contains(strings.ToLower(key), "content-type") && len(value) > 0 {
			return value[0]
		}
	}
	return ""
}
