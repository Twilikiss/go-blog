package common

import (
	"encoding/json"
	"go-blog/log"
	"go-blog/models"
	"net/http"
)

// Success 成功获取并返回数据
func Success(writer http.ResponseWriter, data interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	var result models.Result
	result.Code = 200
	result.Error = ""
	result.Data = data
	resultJson, err := json.Marshal(result)
	if err != nil {
		log.Error("json装配出错:", err)
		Error(writer, err)
		return
	}
	writer.Write(resultJson)
}

func Error(writer http.ResponseWriter, err error) {
	var result models.Result

	result.Code = 900
	result.Error = err.Error()
	result.Data = nil
	errJson, _ := json.Marshal(result)
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(errJson)
}
