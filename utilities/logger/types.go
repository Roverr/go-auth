package logger

// APIPrivateLog is a structure for describing the log message
// to the logger
type APIPrivateLog struct {
	Status   int
	Endpoint string
	UserName string
	Message  string
	ID       uint
	Method   string
}

// APIPublicLog is a structure for describing
// the log message for the logger for public APIs.
type APIPublicLog struct {
	Status   int
	Endpoint string
	Message  string
	Method   string
}
