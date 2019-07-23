package response

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
)

// Error represents the structure of an error message
type Error struct {
	Error   bool   `json:"error"`
	Code    int    `json:"statusCode"`
	Message string `json:"message"`
}

// Errorf return an new error response
func Errorf(w http.ResponseWriter, r *http.Request, err error, code int, message string) {
	logrus.WithFields(logrus.Fields{
		"host":       r.Host,
		"address":    r.RemoteAddr,
		"method":     r.Method,
		"requestURI": r.RequestURI,
		"proto":      r.Proto,
		"useragent":  r.UserAgent(),
	}).WithError(err).Debug(message)

	errorMessage := Error{
		Error:   true,
		Code:    code,
		Message: message,
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(&errorMessage)
	return
}

// Write return a new json response
func Write(w http.ResponseWriter, r *http.Request, data interface{}) {
	logrus.WithFields(logrus.Fields{
		"host":       r.Host,
		"address":    r.RemoteAddr,
		"method":     r.Method,
		"requestURI": r.RequestURI,
		"proto":      r.Proto,
		"useragent":  r.UserAgent(),
	}).Debug(data)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&data)
	return
}
