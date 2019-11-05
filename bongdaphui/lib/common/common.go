package common

// JSON alias type
type JSON = map[string]interface{}

func GenerateResponse(code int, message string, data JSON) JSON {
	if data != nil {
		return JSON{
			"code":    code,
			"message": message,
			"data":    data,
		}
	} else {
		return JSON{
			"code":    code,
			"message": message,
		}
	}
}
