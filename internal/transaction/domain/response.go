package domain

type response struct {
	code    string
	message string
}

func GetApprovedResponse() response {
	return response{
		code:    "00",
		message: "approved",
	}
}

func GetRejectedResponse() response {
	return response{
		code:    "51",
		message: "insufficient founds",
	}
}

func GetGenericResponseError(message string) response {
	return response{
		code:    "07",
		message: message,
	}
}
