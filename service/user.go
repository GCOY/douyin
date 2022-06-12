package service

import (
	"github.com/RaymondCode/simple-demo/dao"
	"github.com/RaymondCode/simple-demo/model"
)

func  GetUserMap(username string) map[string]model.User {
	var usersLoginInfo = map[string]model.User{}
	user := dao.SelectUserByName(username)
	loginInfo := dao.SelectUserInfoByName(username)
	token := loginInfo.UserName + loginInfo.PassWord
	usersLoginInfo[token] = user
	//user := SelectUserByName(name)
	//info := SelectUserInfoByName(name)
	//token := info.UserName + info.PassWord
	//usersLoginInfo[token] = user
	return usersLoginInfo
}

func UserInfoRegister(username string, password string)  {
	dao.UserInfoRegister(username,password)
}

func UserRegister(username string , userid int64)  {
	dao.UserRegister(username,userid)
}

func JudgeUser(username string) bool {
	users := dao.SelectAllUserInfo()
	for _, user := range users {
		if user.UserName == username {
			return false
		}
	}
	return true
}

func GetUserID(username string) int64 {
	return dao.GetUserID(username)
}