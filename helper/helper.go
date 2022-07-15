package helper

import "time"

func ResponseFailedAccess(msg string) map[string]interface{} {
	return map[string]interface{}{
		"code":    "404",
		"message": msg,
	}
}

func ResponseFailedBadRequest(msg string) map[string]interface{} {
	return map[string]interface{}{
		"code":    "400",
		"message": msg,
	}
}

func ResponseFailedServer(msg string) map[string]interface{} {
	return map[string]interface{}{
		"code":    "500",
		"message": msg,
	}
}

func ResponseSuccessNoData(msg string) map[string]interface{} {
	return map[string]interface{}{
		"code":    "200",
		"message": msg,
	}
}

func ResponseSuccessCreate(msg string) map[string]interface{} {
	return map[string]interface{}{
		"code":    "201",
		"message": msg,
	}
}

func ResponseSuccessDelete(msg string) map[string]interface{} {
	return map[string]interface{}{
		"code":    "204",
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
