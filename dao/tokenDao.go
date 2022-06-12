package dao

import (
	"github.com/RaymondCode/simple-demo/connection"
	"github.com/RaymondCode/simple-demo/middleware"
	"github.com/RaymondCode/simple-demo/model"
	"sync"
)

func SaveToken(username string,password string,id int64)  {
	var mutex sync.Mutex
	db := connection.GetDB()
	var token model.Token
	Token := middleware.MD5(username + password)
	token.UserToken = Token
	token.UserID = id
	mutex.Lock()
	tx := db.Begin()
	if err := tx.Create(&token).Error; err != nil {
		tx.Rollback()
		return
	}
	tx.Commit()
	mutex.Unlock()
}
