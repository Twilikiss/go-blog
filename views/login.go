package views

import (
	"go-blog/common"
	"go-blog/config"
	"go-blog/log"
	"net/http"
)

func (*htmlApi) Login(writer http.ResponseWriter, request *http.Request) {
	t := common.Template
	err := t.ExecuteHTML(writer, "login.html", config.Cfg.Viewer)
	if err != nil {
		log.Errorf("index.html write data error:%s", err)
		t.WriteError(writer, "登录页面显示异常，请联系管理员！~")
		return
	}
}
