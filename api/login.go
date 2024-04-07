package api

import (
	"go-blog/common"
	"go-blog/service"
	"net/http"
)

func (*apiHandler) Login(writer http.ResponseWriter, request *http.Request) {
	// 接收用户端传入的用户名和密码(json)
	params := common.GetRequestJsonParam(request)
	userName := params["username"]
	passwd := params["passwd"]

	data, err := service.Login(userName, passwd)
	if err != nil {
		common.Error(writer, err)
		return
	}
	common.Success(writer, data)
}
