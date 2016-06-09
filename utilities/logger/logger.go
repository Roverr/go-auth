package logger

import (
	"os"

	"github.com/op/go-logging"
)

// Standard is a logger is the basic instant good-to-go
// logger of the project.
var Standard = logging.MustGetLogger("standard")

// InitLogger is initializing the logger
// for the webserver.
func InitLogger() {
	standardLogFormat := logging.MustStringFormatter(
		`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
	)
	standardBackend := logging.NewLogBackend(os.Stderr, "", 0)
	standardBackendFormatter := logging.NewBackendFormatter(standardBackend, standardLogFormat)

	logging.SetBackend(standardBackendFormatter)
}

// PublicAPIMessage is a function for formatting a logger message
// which can be called when a non-authenticated endpoint is reached.
func PublicAPIMessage(data APIPublicLog) {
	Standard.Infof("%d - %s %s - User: Unknown - %s ", data.Status, data.Method, data.Endpoint, data.Message)
}

// PrivateAPIMessage is a function for formatting a logger message
// which can be called from authenticated endpoints for correct logging.
func PrivateAPIMessage(data APIPrivateLog) {
	Standard.Infof("%d - %s %s - User: %s ID: %d - %s", data.Status, data.Method, data.Endpoint, data.UserName, data.ID, data.Message)
}
