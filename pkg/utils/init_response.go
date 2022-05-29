package utils

func SuccessResponse(resp map[string]any) map[string]any {
	resp["status_code"] = 0
	resp["status_msg"] = "OK"
	return resp
}

func FailResponse(msg string) map[string]any {
	return map[string]any{
		"status_code": 1,
		"status_msg":  msg,
	}
}
