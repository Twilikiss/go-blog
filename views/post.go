package views

import (
	"errors"
	"go-blog/common"
	"go-blog/log"
	"go-blog/service"
	"go-blog/utils"
	"net/http"
	"strconv"
	"strings"
)

func (*htmlApi) Detail(writer http.ResponseWriter, request *http.Request) {
	t := common.Template

	// =======对IP进行检查========
	ip := utils.GetClientIp(request)
	if ip != "" {
		// 限制同一ip恶意刷观看浏览量
		client := utils.NewClient()
		err := utils.CountString(client, ip)
		if err != nil {
			log.Error("user access err", err)
			common.Error(writer, errors.New("暂时访问，稍后重试看看吧！~"))
			return
		}
	}

	// 获取到path中携带的具体文章id
	path := request.URL.Path
	log.Info("get the path:", path)
	pidStr := strings.TrimPrefix(path, "/p/")
	pidF := strings.TrimSuffix(pidStr, ".html")
	pid, err := strconv.Atoi(pidF)
	if err != nil {
		log.Error("category disable to get the param from path:", err)
		t.WriteError(writer, "参数错误，请检查访问链接！~")
		return
	}
	postRes, err := service.GetPostDetailByPid(pid)
	if err != nil {
		log.Error("detail.html cannot find the specified data:", err)
		t.WriteError(writer, "无法找到文章详情数据，请重试")
		return
	}
	err = t.ExecuteHTML(writer, "detail.html", postRes)
	if err != nil {
		log.Errorf("index.html write data error:%s", err)
		t.WriteError(writer, "登录页面显示异常，请联系管理员！~")
		return
	}
}
