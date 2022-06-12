package dao

import (
	"github.com/RaymondCode/simple-demo/connection"
	"github.com/RaymondCode/simple-demo/model"
	"sync"
)

//查询login_infos表所有用户信息
func SelectAllUserInfo() []model.LoginInfo {
	db := connection.GetDB()
	//var user User
	var loginInfos []model.LoginInfo
	db.Find(&loginInfos)
	return loginInfos
}

func SelectAllUser() []model.User {
	db := connection.GetDB()
	var users []model.User
	db.Find(&users)
	return users
}

//向login_infos表中插入信息
func UserInfoRegister(username string, password string)  {
	var mutex sync.Mutex
	db := connection.GetDB()
	mutex.Lock()
	tx := db.Begin()
	login := model.LoginInfo{UserName: username,PassWord: password}
	if err := tx.Create(&login).Error ;err != nil {
		tx.Rollback()
		return
	}
	tx.Commit()
	mutex.Unlock()
}

//向users表中插入信息
func UserRegister(username string , userid int64)  {
	var mutex sync.Mutex
	db := connection.GetDB()
	mutex.Lock()
	tx := db.Begin()
	user := model.User{Id: userid,Name: username}
	if err := db.Create(&user).Error; err != nil {
		tx.Rollback()
		return
	}
	tx.Commit()
	mutex.Unlock()
}

//根据username向login_info表中查询id
func GetUserID(username string) int64 {
	db := connection.GetDB()
	var login model.LoginInfo
	db.Where("user_name=?", username).Find(&login)
	return login.ID
}

//根据username向users表中查询用户
func SelectUserByName(username string) model.User {
	db := connection.GetDB()
	var user model.User
	db.Where("name = ?",username).Find(&user)
	return user
}

//根据username向login_infos表中查询用户信息
func SelectUserInfoByName(username string) model.LoginInfo {
	db := connection.GetDB()
	var loginInfo model.LoginInfo
	db.Where("user_name = ?",username).Find(&loginInfo)
	return loginInfo
}

//获取token，调用SelectUserIdByToken获取id，根据id查用户
func SelectUserByToken(token string) model.User {
	db := connection.GetDB()
	id := SelectUserIdByToken(token)
	var user model.User
	db.Where("id = ?",id).Find(&user)
	return user
}

//在tokens表中根据token查用户id
func SelectUserIdByToken(token string) int64 {
	db := connection.GetDB()
	var UserToken model.Token
	db.Where("user_token = ?",token).Find(&UserToken)
	return UserToken.UserID
}

