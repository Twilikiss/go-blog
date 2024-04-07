package dao

import (
	"go-blog/log"
	"go-blog/models"
	"go-blog/utils"
)

var client = utils.NewClient()

func GetAllPostPage(page, pageSize int) ([]models.Post, error) {
	session := MysqlEngine.NewSession()
	page = (page - 1) * pageSize
	rows, err := session.Raw("select * from blog_post limit ?,?", page, pageSize).QueryRows()
	if err != nil {
		log.Errorf("post search is error: %s", err)
		return nil, err
	}
	var posts []models.Post
	for rows.Next() {
		var post models.Post
		err := rows.Scan(
			&post.Pid,
			&post.Title,
			&post.Content,
			&post.Markdown,
			&post.CategoryId,
			&post.UserId,
			&post.ViewCount,
			&post.Type,
			&post.Slug,
			&post.CreateAt,
			&post.UpdateAt,
		)
		if err != nil {
			log.Errorf("blog_post读取数据出错：%s", err)
			return nil, err
		}
		// 暂时改为使用redis来记录点击量
		post.ViewCount = utils.GetViews(client, post.Pid)
		posts = append(posts, post)
	}
	return posts, err
}

func CountAllPost() int {
	session := MysqlEngine.NewSession()
	row := session.Raw("select count(*) from blog_post").QueryRow()
	if row.Err() != nil {
		log.Errorf("CountAllPost is error: %s", row.Err())
	}
	var count int
	_ = row.Scan(&count)
	return count
}

func GetPostPageByCid(cId, page, pageSize int) ([]models.Post, error) {
	session := MysqlEngine.NewSession()
	page = (page - 1) * pageSize
	rows, err := session.Raw("select * from blog_post where category_id=? limit ?,?", cId, page, pageSize).QueryRows()
	if err != nil {
		log.Errorf("post search is error: %s", err)
		return nil, err
	}
	var posts []models.Post
	for rows.Next() {
		var post models.Post
		err := rows.Scan(
			&post.Pid,
			&post.Title,
			&post.Content,
			&post.Markdown,
			&post.CategoryId,
			&post.UserId,
			&post.ViewCount,
			&post.Type,
			&post.Slug,
			&post.CreateAt,
			&post.UpdateAt,
		)
		if err != nil {
			log.Errorf("blog_post读取数据出错：%s", err)
			return nil, err
		}
		// 暂时改为使用redis来记录点击量
		post.ViewCount = utils.GetViews(client, post.Pid)
		posts = append(posts, post)
	}
	return posts, err
}

func CountPostByCid(cid int) int {
	session := MysqlEngine.NewSession()
	row := session.Raw("select count(*) from blog_post where category_id = ?", cid).QueryRow()
	if row.Err() != nil {
		log.Errorf("CountAllPost is error: %s", row.Err())
	}
	var count int
	_ = row.Scan(&count)
	return count
}

func GetPostByPid(pid int) (*models.Post, error) {
	session := MysqlEngine.NewSession()
	row := session.Raw("select * from blog_post where pid = ? limit 1", pid).QueryRow()
	if row.Err() != nil {
		log.Errorf("GetPostByPid is error: %s", row.Err())
		return nil, row.Err()
	}
	var post models.Post
	err := row.Scan(
		&post.Pid,
		&post.Title,
		&post.Content,
		&post.Markdown,
		&post.CategoryId,
		&post.UserId,
		&post.ViewCount,
		&post.Type,
		&post.Slug,
		&post.CreateAt,
		&post.UpdateAt,
	)
	if err != nil {
		log.Errorf("blog_post读取数据出错：%s", err)
		return nil, err
	}
	// 暂时改为使用redis来记录点击量
	go utils.AddViews(client, post.Pid)
	post.ViewCount = utils.GetViews(client, post.Pid)
	return &post, nil
}

func SavePost(post *models.Post) (*models.Post, error) {
	session := MysqlEngine.NewSession()
	result, err := session.Raw("insert into blog_post "+
		"(title,content,markdown,category_id,user_id,view_count,type,slug,create_at,update_at) "+
		"values(?,?,?,?,?,?,?,?,?,?)",
		post.Title,
		post.Content,
		post.Markdown,
		post.CategoryId,
		post.UserId,
		post.ViewCount,
		post.Type,
		post.Slug,
		post.CreateAt,
		post.UpdateAt,
	).Exec()
	pid, err := result.LastInsertId()
	post.Pid = int(pid)
	if err != nil {
		log.Errorf("blog_post保存文章数据出错：%s", err)
		return nil, err
	}
	return post, nil
}

func UpdatePost(post *models.Post) error {
	session := MysqlEngine.NewSession()
	_, err := session.Raw("update blog_post "+
		"set title=?,content=?,markdown=?,category_id=?,view_count=?,type=?,slug=?,update_at=? "+
		"where pid=?",
		post.Title,
		post.Content,
		post.Markdown,
		post.CategoryId,
		post.ViewCount,
		post.Type,
		post.Slug,
		post.UpdateAt,
		post.Pid,
	).Exec()
	if err != nil {
		log.Errorf("blog_post修改文章出错：%s", err)
		return err
	}
	return nil
}

func GetAllPost() ([]models.Post, error) {
	session := MysqlEngine.NewSession()
	rows, err := session.Raw("select * from blog_post").QueryRows()
	if err != nil {
		log.Errorf("post search is error: %s", err)
		return nil, err
	}
	var posts []models.Post
	for rows.Next() {
		var post models.Post
		err := rows.Scan(
			&post.Pid,
			&post.Title,
			&post.Content,
			&post.Markdown,
			&post.CategoryId,
			&post.UserId,
			&post.ViewCount,
			&post.Type,
			&post.Slug,
			&post.CreateAt,
			&post.UpdateAt,
		)
		if err != nil {
			log.Errorf("blog_post读取数据出错：%s", err)
			return nil, err
		}
		// 暂时改为使用redis来记录点击量
		post.ViewCount = utils.GetViews(client, post.Pid)
		posts = append(posts, post)
	}
	return posts, err
}

func GetAllPostPageBySlug(slug string, page, pageSize int) ([]models.Post, error) {
	session := MysqlEngine.NewSession()
	page = (page - 1) * pageSize
	rows, err := session.Raw("select * from blog_post where slug=? limit ?,?", slug, page, pageSize).QueryRows()
	if err != nil {
		log.Errorf("post search is error: %s", err)
		return nil, err
	}
	var posts []models.Post
	for rows.Next() {
		var post models.Post
		err := rows.Scan(
			&post.Pid,
			&post.Title,
			&post.Content,
			&post.Markdown,
			&post.CategoryId,
			&post.UserId,
			&post.ViewCount,
			&post.Type,
			&post.Slug,
			&post.CreateAt,
			&post.UpdateAt,
		)
		if err != nil {
			log.Errorf("blog_post读取数据出错：%s", err)
			return nil, err
		}
		// 暂时改为使用redis来记录点击量
		post.ViewCount = utils.GetViews(client, post.Pid)
		posts = append(posts, post)
	}
	return posts, err
}

func CountPostBySlug(slug string) int {
	session := MysqlEngine.NewSession()
	row := session.Raw("select count(*) from blog_post where slug = ?", slug).QueryRow()
	if row.Err() != nil {
		log.Errorf("CountAllPost is error: %s", row.Err())
	}
	var count int
	_ = row.Scan(&count)
	return count
}

func GetPostSearch(searchVar string) ([]models.SearchResp, error) {
	session := MysqlEngine.NewSession()
	rows, err := session.Raw("select pid, title from blog_post where title like ?", "%"+searchVar+"%").QueryRows()
	if err != nil {
		log.Errorf("post search is error: %s", err)
		return nil, err
	}
	var searchResps []models.SearchResp
	for rows.Next() {
		var searchResp models.SearchResp
		err = rows.Scan(&searchResp.Pid, &searchResp.Title)
		if err != nil {
			log.Errorf("blog_post读取数据出错：%s", err)
			return nil, err
		}
		searchResps = append(searchResps, searchResp)
	}
	return searchResps, nil
}
