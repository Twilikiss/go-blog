package service

import (
	"go-blog/config"
	"go-blog/dao"
	"go-blog/log"
	"go-blog/models"
	"go-blog/utils"
	"html/template"
)

func GetIndexInfoByCid(cid, page, pageSize int) (*models.CategoryResponse, error) {
	categories, err := dao.GetAllCategory()
	if err != nil {
		log.Errorf("Get IndexInfo error:%s", err)
		return nil, err
	}

	posts, err := dao.GetPostPageByCid(cid, page, pageSize)
	var postMores []models.PostMore
	categoryName := dao.GetCategoryNameById(cid)
	for _, post := range posts {
		userName := dao.GetUserNameById(post.UserId)

		//content := []rune(post.Content)
		//if len(content) > 100 {
		//	content = content[:100]
		//}

		content := utils.TrimHtml(post.Content)

		if len(content) >= 100 {
			content = string([]rune(utils.TrimHtml(post.Content))[:100])
		}

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

	total := dao.CountPostByCid(cid)
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
	categoryResponse := &models.CategoryResponse{
		HomeResponse: hr,
		CategoryName: categoryName,
	}
	return categoryResponse, nil
}
