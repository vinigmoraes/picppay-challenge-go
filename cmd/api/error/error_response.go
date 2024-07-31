package errorhandler

type ErrorResponse struct {
	FieldName string
	Message   string
	Value     interface{}
	Error     bool
}
