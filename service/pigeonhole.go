package service

import (
	"go-blog/config"
	"go-blog/dao"
	"go-blog/models"
)

func GetPostPigeonhole() models.PigeonholeRes {
	// 查询所有的文章，进行月份整理
	posts, _ := dao.GetAllPost()
	// 查询所有的文章分类
	categories, _ := dao.GetAllCategory()

	lines := make(map[string][]models.Post)

	for _, post := range posts {
		at := post.CreateAt
		month := at.Format("2006-01")
		lines[month] = append(lines[month], post)
	}

	return models.PigeonholeRes{
		Viewer:       config.Cfg.Viewer,
		SystemConfig: config.Cfg.System,
		Categories:   categories,
		Lines:        lines,
	}
}
