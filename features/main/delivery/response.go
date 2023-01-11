package delivery

func SuccessResponseData(data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"data": data,
	}
}

func SuccessResponseNoData(msg string) map[string]interface{} {
	return map[string]interface{}{
		"message": msg,
	}
}

func FailResponse(msg interface{}) map[string]interface{} {
	return map[string]interface{}{
		"message": msg,
	}
}
