package domain

type Response struct {
	Code    string
	Message string
}

func GetApprovedResponse() Response {
	return Response{
		Code:    "00",
		Message: "approved",
	}
}

func GetRejectedResponse() Response {
	return Response{
		Code:    "51",
		Message: "insufficient funds",
	}
}

func GetGenericResponseError(message string) Response {
	return Response{
		Code:    "07",
		Message: message,
	}
}
