package dao

import (
	"go-blog/log"
	"go-blog/models"
)

func GetUserNameById(uid int) string {
	session := MysqlEngine.NewSession()
	row := session.Raw("select user_name from blog_user where uid=?", uid).QueryRow()
	if row.Err() != nil {
		log.Errorf("GetUserNameById is error: %s", row.Err())
	}
	var name string
	_ = row.Scan(&name)
	return name
}

func GetUserByInfo(userName, passwd string) *models.UserInfo {
	session := MysqlEngine.NewSession()
	row := session.Raw("select uid,user_name,avatar from blog_user where user_name=? and passwd=? limit 1", userName, passwd).QueryRow()
	if row.Err() != nil {
		log.Errorf("GetUserByInfo is error: %s", row.Err())
		return nil
	}
	var userInfo models.UserInfo
	err := row.Scan(&userInfo.Uid, &userInfo.UserName, &userInfo.Avatar)
	if err != nil {
		log.Info("rows.Scan is error:", err)
		return nil
	}
	return &userInfo
}
