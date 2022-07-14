package helper

import "time"

func ResponseFailed(msg string) map[string]interface{} {
	return map[string]interface{}{
		"code":    "404",
		"message": msg,
	}
}

func ResponseSuccessNoData(msg string) map[string]interface{} {
	return map[string]interface{}{
		"code":    "200",
		"message": msg,
	}
}

func ResponseSuccessWithData(msg string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"code":        "200",
		"message":     msg,
		"currenttime": time.Now().Format("2006-01-02 15:04:05"),
		"data":        data,
	}
}

func ResponseSuccessWithDataPage(msg string, page int64, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"code":        "200",
		"message":     msg,
		"currenttime": time.Now().Format("2006-01-02 15:04:05"),
		"totalpage":   page,
		"data":        data,
	}
}
