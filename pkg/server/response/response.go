package response

import (
	"encoding/json"
	"mime"
	"net/http"

	"github.com/erik-sostenes/notifications-api/pkg/domain/wrongs"
)

// Response represents the http request body
type Response struct {
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
}

// JSON encodes the body of the http request in json format
func JSON(w http.ResponseWriter, method int, body any) error {
	w.Header().Add("Content-type", "application/json; charset=utf-8")
	w.WriteHeader(method)

	return json.NewEncoder(w).Encode(body)
}

// Bind parses the request data to an object
// checks the content type
func Bind(w http.ResponseWriter, r *http.Request, body any) (ok bool, err error) {
	content := r.Header.Get("Content-type")

	if content == "" {
		_ = JSON(w, http.StatusUnprocessableEntity, Response{
			Message: "missing content type",
		})
		return
	}

	mediaType, _, err := mime.ParseMediaType(content)
	if err != nil {
		_ = JSON(w, http.StatusUnprocessableEntity, Response{
			Message: err.Error(),
		})
		return
	}

	switch mediaType {
	case "application/json; charset=utf-8", "application/json":
		err = json.NewDecoder(r.Body).Decode(body)
		if err != nil {
			_ = JSON(w, http.StatusUnprocessableEntity, Response{
				Message: "the format of the body of the request is malformed",
			})
		}
		return true, nil
	default:
		_ = JSON(w, http.StatusUnsupportedMediaType, Response{
			Message: "unsupported media type",
		})
		return
	}
}

// ErrorHandler handles http error response depending on error type
func ErrorHandler(w http.ResponseWriter, err error) error {
	switch err.(type) {
	case wrongs.StatusBadRequest:
		return JSON(w, http.StatusBadRequest, Response{Message: err.Error()})
	case wrongs.StatusUnprocessableEntity:
		return JSON(w, http.StatusUnprocessableEntity, Response{Message: err.Error()})
	case wrongs.StatusNotFound:
		return JSON(w, http.StatusNotFound, Response{Message: err.Error()})
	default:
		return JSON(w, http.StatusInternalServerError, Response{Message: err.Error()})
	}
}
