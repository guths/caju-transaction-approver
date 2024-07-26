package domain

type Response struct {
	code    string
	message string
}

func GetApprovedResponse() Response {
	return Response{
		code:    "00",
		message: "approved",
	}
}

func GetRejectedResponse() Response {
	return Response{
		code:    "51",
		message: "insufficient funds",
	}
}

func GetGenericResponseError(message string) Response {
	return Response{
		code:    "07",
		message: message,
	}
}
