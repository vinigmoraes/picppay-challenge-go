package errorhandler

func getStatusCodeByError(apiError APIError) int {
	var statusCode []map[APIError]int
	userAlreadyExists := map[APIError]int{apiError.GetErrorType(): 409}

	statusCodePopulatedList := append(statusCode, userAlreadyExists)

	return statusCodePopulatedList[0][apiError.GetErrorType()]
}

func HandleError(err APIError) (string, int) {
	var message string
	var code int

	message = err.GetError()
	code = getStatusCodeByError(err)

	return message, code
}
