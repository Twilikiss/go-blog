package views

import (
	"go-blog/common"
	"go-blog/log"
	"net/http"
)

func (*htmlApi) About(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "text/html")
	writer.WriteHeader(200)

	t := common.Template

	err := t.ExecuteHTML(writer, "about.html", nil)
	if err != nil {
		log.Errorf("index.html write data error:%s", err)
		t.WriteError(writer, "无法读取数据，请联系管理员！~")
		return
	}
}
