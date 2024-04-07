package dao

import (
	"go-blog/log"
	"go-blog/models"
)

func GetAllCategory() ([]models.Category, error) {
	session := MysqlEngine.NewSession()
	rows, err := session.Raw("select * from blog_category").QueryRows()
	if err != nil {
		log.Errorf("category search is error: %s", err)
		return nil, err
	}
	var data []models.Category
	for rows.Next() {
		var category models.Category
		err = rows.Scan(&category.Cid, &category.Name, &category.CreateAt, &category.UpdateAt)
		if err != nil {
			log.Errorf("blog_category读取数据出错：%s", err)
			return nil, err
		}
		data = append(data, category)
	}
	return data, nil
}

func GetCategoryNameById(cId int) string {
	session := MysqlEngine.NewSession()
	row := session.Raw("select name from blog_category where cid=?", cId).QueryRow()
	if row.Err() != nil {
		log.Errorf("GetCategoryNameById is error: %s", row.Err())
	}
	var name string
	_ = row.Scan(&name)
	return name
}
