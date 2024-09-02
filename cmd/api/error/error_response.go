package errorhandler

type ErrorResponse struct {
	FieldName string      `json:"field_name"`
	Message   string      `json:"message"`
	Value     interface{} `json:"value"`
	Error     bool        `json:"error"`
}
