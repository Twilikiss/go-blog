package service

import (
	"go-blog/config"
	"go-blog/dao"
	"go-blog/log"
	"go-blog/models"
	"go-blog/utils"
	"html/template"
)

func GetAllIndexInfo(slug string, page, pageSize int) (*models.HomeResponse, error) {
	categories, err := dao.GetAllCategory()
	if err != nil {
		log.Errorf("Get IndexInfo error:%s", err)
		return nil, err
	}
	var posts []models.Post
	var flag = false // 是否启动查询所有结果
	if slug == "" {
		posts, err = dao.GetAllPostPage(page, pageSize)
		flag = true
	} else {
		posts, err = dao.GetAllPostPageBySlug(slug, page, pageSize)
	}
	if len(posts) == 0 {
		posts, err = dao.GetAllPostPage(page, pageSize)
		flag = true
	}
	var postMores []models.PostMore
	for _, post := range posts {
		categoryName := dao.GetCategoryNameById(post.CategoryId)
		userName := dao.GetUserNameById(post.UserId)

		content := utils.TrimHtml(post.Content)

		if len(content) >= 100 {
			content = string([]rune(utils.TrimHtml(post.Content))[:100])
		}

		//content := []rune(post.Content)
		//if len(content) > 100 {
		//	content = content[0:100]
		//}
		postMore := models.PostMore{
			Pid:          post.Pid,
			Title:        post.Title,
			Slug:         post.Slug,
			Content:      template.HTML(content),
			CategoryId:   post.CategoryId,
			CategoryName: categoryName,
			UserId:       post.UserId,
			UserName:     userName,
			ViewCount:    post.ViewCount,
			Type:         post.Type,
			CreateAt:     post.CreateAt.Format("2006-01-02 15:04:05"),
			UpdateAt:     post.UpdateAt.Format("2006-01-02 15:04:05"),
		}
		postMores = append(postMores, postMore)
	}

	var total int

	if flag {
		total = dao.CountAllPost()
	} else {
		total = dao.CountPostBySlug(slug)
	}

	pagesCount := (total-1)/10 + 1
	var pages []int
	for i := 1; i <= pagesCount; i++ {
		pages = append(pages, i)
	}

	var hr = &models.HomeResponse{
		Viewer:     config.Cfg.Viewer,
		Categories: categories,
		Posts:      postMores,
		Total:      total,
		Page:       page,
		Pages:      pages,
		PageEnd:    page != pagesCount,
	}
	return hr, nil
}
