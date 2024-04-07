package views

import (
	"go-blog/common"
	"go-blog/log"
	"go-blog/service"
	"net/http"
)

func (*htmlApi) Pigeonhole(writer http.ResponseWriter, request *http.Request) {
	t := common.Template

	pigeonholes := service.GetPostPigeonhole()
	err := t.ExecuteHTML(writer, "pigeonhole.html", pigeonholes)
	if err != nil {
		log.Errorf("index.html write data error:%s", err)
		t.WriteError(writer, "登录页面显示异常，请联系管理员！~")
		return
	}
}
