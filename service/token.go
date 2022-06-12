package service

import "github.com/RaymondCode/simple-demo/dao"

func SaveToken(username string,password string,id int64)  {
	dao.SaveToken(username,password,id)
}
