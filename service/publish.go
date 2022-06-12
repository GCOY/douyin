package service

import (
	"github.com/RaymondCode/simple-demo/dao"
	"github.com/RaymondCode/simple-demo/model"
)

func GetPublishList(id int64) []model.Video {
	return dao.GetPublishList(id)
}

