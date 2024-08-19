package errorhandler

type APIError interface {
	GetError() string
	GetErrorType() APIError
}
