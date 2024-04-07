package views

import (
	"go-blog/common"
	"go-blog/log"
	"go-blog/service"
	"net/http"
)

func (*htmlApi) Writing(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "text/html")
	writer.WriteHeader(200)

	t := common.Template

	wr, err := service.Writing()

	err = t.ExecuteHTML(writer, "writing.html", wr)
	if err != nil {
		log.Errorf("writing.html write data error:%s", err)
		t.WriteError(writer, "无法读取数据，请联系管理员！~")
		return
	}
}
