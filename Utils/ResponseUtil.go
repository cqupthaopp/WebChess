package Utils

import "github.com/gin-gonic/gin"

func GetNormalRes(data map[string]interface{}) map[string]interface{} {

	return gin.H{
		"state": 10001,
		"info":  "success",
		"data":  data,
	}

}

func GetErrorRes(data map[string]interface{}) map[string]interface{} {

	return gin.H{
		"state": 10002,
		"info":  "error",
		"data":  data,
	}

}
