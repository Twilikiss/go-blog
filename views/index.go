package views

import (
	"go-blog/common"
	"go-blog/log"
	"go-blog/service"
	"net/http"
	"strconv"
	"strings"
)

type IndexData struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

func (html *htmlApi) Index(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "text/html")
	writer.WriteHeader(200)

	t := common.Template

	// 数据库查询
	if err := request.ParseForm(); err != nil {
		log.Error("表单解析出错:", err)
		t.WriteError(writer, "表单解析出错，请联系管理员！~")
		return
	}
	// 获取当前的页码
	pageStr := request.Form.Get("page")
	page := 1
	if pageStr != "" {
		page, _ = strconv.Atoi(pageStr)
	}
	// 获取每页显示的文章数量
	pageSize := 10

	path := request.URL.Path
	slug := strings.TrimPrefix(path, "/")

	hr, err := service.GetAllIndexInfo(slug, page, pageSize)
	if err != nil {
		log.Errorf("index.html get data error:%s", err)
		t.WriteError(writer, "数据库出错，请联系管理员！~")
		return
	}
	err = t.ExecuteHTML(writer, "index.html", hr)
	if err != nil {
		log.Errorf("index.html write data error:%s", err)
		t.WriteError(writer, "无法读取数据，请联系管理员！~")
		return
	}
}
