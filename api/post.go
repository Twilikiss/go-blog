package api

import (
	"errors"
	"go-blog/common"
	"go-blog/dao"
	"go-blog/log"
	"go-blog/models"
	"go-blog/service"
	"go-blog/utils"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func (*apiHandler) SaveAndUpdatePost(writer http.ResponseWriter, request *http.Request) {
	// 先要判断一下操作的类型
	m := request.Method
	switch m {
	case http.MethodPost:
		// 如果是上传文章，接着就要获取参数
		params := common.GetRequestJsonParam(request)
		categoryId, _ := strconv.Atoi(params["categoryId"].(string))
		content := params["content"].(string)
		markdown := params["markdown"].(string)
		slug := params["slug"].(string)
		title := params["title"].(string)
		postType, ok := params["type"].(int)
		if !ok {
			postType = 0
		}

		// 获取用户id
		token := request.Header.Get("Authorization")
		_, claim, err := utils.ParseToken(token)
		if err != nil {
			log.Errorf("token error:%v", err)
			common.Error(writer, errors.New("登录已过期"))
			return
		}
		uid := claim.Uid
		var post = &models.Post{
			Title:      title,
			Slug:       slug,
			Content:    content,
			Markdown:   markdown,
			CategoryId: categoryId,
			UserId:     uid,
			ViewCount:  0,
			Type:       postType,
			CreateAt:   time.Now(),
			UpdateAt:   time.Now(),
		}

		post, err = service.SavePost(post)
		if err != nil {
			common.Error(writer, err)
			return
		}
		common.Success(writer, post)
	case http.MethodPut:
		// 编辑更新操作
		params := common.GetRequestJsonParam(request)
		categoryId, _ := strconv.Atoi(params["categoryId"].(string))
		content := params["content"].(string)
		markdown := params["markdown"].(string)
		slug := params["slug"].(string)
		title := params["title"].(string)
		pid := int(params["pid"].(float64))
		postType, ok := params["type"].(int)
		if !ok {
			postType = 0
		}
		// 获取用户id
		token := request.Header.Get("Authorization")
		_, claim, err := utils.ParseToken(token)
		if err != nil {
			log.Errorf("token error:%v", err)
			common.Error(writer, errors.New("登录已过期"))
			return
		}
		uid := claim.Uid
		var post = &models.Post{
			Pid:        pid,
			Title:      title,
			Slug:       slug,
			Content:    content,
			Markdown:   markdown,
			CategoryId: categoryId,
			UserId:     uid,
			ViewCount:  0,
			Type:       postType,
			CreateAt:   time.Now(),
			UpdateAt:   time.Now(),
		}

		err = service.UpdatePost(post)
		if err != nil {
			common.Error(writer, errors.New("保存修改失败"))
			return
		}
		common.Success(writer, post)
	}
}

func (*apiHandler) GetPost(writer http.ResponseWriter, request *http.Request) {

	path := request.URL.Path
	log.Info("get the path:", path)
	pidStr := strings.TrimPrefix(path, "/api/v1/post/")
	pid, err := strconv.Atoi(pidStr)
	if err != nil {
		log.Error("category disable to get the param from path:", err)
		common.Error(writer, errors.New("参数错误，请检查访问链接！~"))
		return
	}
	post, err := dao.GetPostByPid(pid)
	if err != nil {
		log.Error("detail.html cannot find the specified data:", err)
		common.Error(writer, errors.New("无法找到文章详情数据，请重试"))
		return
	}
	common.Success(writer, post)
}

func (*apiHandler) SearchPost(writer http.ResponseWriter, request *http.Request) {
	// 数据库查询
	if err := request.ParseForm(); err != nil {
		log.Error("表单解析出错:", err)
		common.Error(writer, errors.New("表单解析出错，请联系管理员！~"))
		return
	}
	// 获取当前的页码
	searchVar := request.Form.Get("val")

	searchResp, err := service.SearchPost(searchVar)
	if err != nil {
		log.Error("模糊查询错误:", err)
		common.Error(writer, errors.New("无法找到文章详情数据，请重试"))
		return
	}
	common.Success(writer, searchResp)
}
