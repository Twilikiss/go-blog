package service

import (
	"go-blog/config"
	"go-blog/dao"
	"go-blog/log"
	"go-blog/models"
	"html/template"
)

func GetPostDetailByPid(pid int) (*models.PostRes, error) {
	post, err := dao.GetPostByPid(pid)
	if err != nil {
		log.Error("Get Post error:", err)
		return nil, err
	}

	categoryName := dao.GetCategoryNameById(post.CategoryId)
	userName := dao.GetUserNameById(post.UserId)

	postMore := models.PostMore{
		Pid:          post.Pid,
		Title:        post.Title,
		Slug:         post.Slug,
		Content:      template.HTML(post.Content),
		CategoryId:   post.CategoryId,
		CategoryName: categoryName,
		UserId:       post.UserId,
		UserName:     userName,
		ViewCount:    post.ViewCount,
		Type:         post.Type,
		CreateAt:     post.CreateAt.Format("2006-01-02 15:04:05"),
		UpdateAt:     post.UpdateAt.Format("2006-01-02 15:04:05"),
	}

	var postRes = &models.PostRes{
		Viewer:       config.Cfg.Viewer,
		SystemConfig: config.Cfg.System,
		Article:      postMore,
	}

	return postRes, nil
}

func Writing() (*models.WritingRes, error) {
	var wr models.WritingRes
	wr.Title = config.Cfg.Viewer.Title
	wr.CdnURL = config.Cfg.System.CdnURL
	category, err := dao.GetAllCategory()
	if err != nil {
		log.Errorf("GetAllCategory is error:", err)
		return nil, err
	}
	wr.Categories = category
	return &wr, nil
}

func SavePost(post *models.Post) (*models.Post, error) {
	return dao.SavePost(post)
}

func UpdatePost(post *models.Post) error {
	return dao.UpdatePost(post)
}

func SearchPost(searchVar string) ([]models.SearchResp, error) {
	return dao.GetPostSearch(searchVar)
}
