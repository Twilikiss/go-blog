package router

import (
	"go-blog/api"
	"go-blog/views"
	"net/http"
)

func Router() {
	// 导入静态资源
	http.Handle("/resource/", http.StripPrefix("/resource/", http.FileServer(http.Dir("public/resource/"))))

	// 1.页面 views 2.数据（json） api
	http.HandleFunc("/head", api.API.TestHeader)
	http.HandleFunc("/index", api.API.TestIndex)
	http.HandleFunc("/api/v1/login", api.API.Login)
	http.HandleFunc("/api/v1/post", api.API.SaveAndUpdatePost)
	http.HandleFunc("/api/v1/post/", api.API.GetPost)
	http.HandleFunc("/api/v1/qiniu/token", api.API.QiniuToken)
	http.HandleFunc("/api/v1/post/search", api.API.SearchPost)

	// 下面是返回的前端页面的相关路由
	// 首页的路由
	http.HandleFunc("/", views.HTML.Index)

	// http://localhost:8080/c/1
	// 这里我们可以仿照我们在7days项目中实现的web框架相关内容进行改写
	http.HandleFunc("/c/", views.HTML.Category)

	// 登录页面的路由
	http.HandleFunc("/login", views.HTML.Login)

	// http://localhost:8080/p/7.html
	// 文章详情页面
	http.HandleFunc("/p/", views.HTML.Detail)

	// 写作页面的路由
	http.HandleFunc("/writing", views.HTML.Writing)

	// 归档页面
	http.HandleFunc("/pigeonhole", views.HTML.Pigeonhole)

	// "关于"页面
	http.HandleFunc("/about", views.HTML.About)
}
