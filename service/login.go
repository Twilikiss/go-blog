package service

import (
	"errors"
	"fmt"
	"go-blog/dao"
	"go-blog/log"
	"go-blog/models"
	"go-blog/utils"
)

func Login(username, passwd interface{}) (*models.LoginResponse, error) {
	// TODO 注意go语言中的类型转换相关操作，注意做笔记！！！！
	usernameString := fmt.Sprint(username)
	passwdString := fmt.Sprint(passwd)

	// 先对密码进行md5加密，在比对数据库中存有的加密后的数据
	passwdString = utils.Md5Crypt(passwdString, "elysia")

	userInfo := dao.GetUserByInfo(usernameString, passwdString)
	if userInfo == nil {
		log.Infof("Failed to find user:%v", usernameString)
		return nil, errors.New("用户名或者密码错误，请重试")
	}
	// 下面要构建我们的Token
	token, err := utils.Award(userInfo.Uid)
	if err != nil {
		log.Errorf("Failed to create token:%v", err)
		return nil, errors.New("后台登录处理错误，请重试")
	}

	var lr = &models.LoginResponse{
		Token:    token,
		UserInfo: userInfo,
	}

	return lr, nil
}
