package utils

func BuildResponse(err error) Response {

	var response Response

	if err != nil {
		errCode, errMessage := ErrorCodeAndMessage(err)
		response.StatusCode = errCode
		response.Messages = errMessage
		response.Data = nil
		return response
	}

	response.StatusCode = 200
	response.Messages = "Success"
	response.Data = response

	return response

}
