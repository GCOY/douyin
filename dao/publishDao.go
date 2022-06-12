package dao

import (
	"github.com/RaymondCode/simple-demo/connection"
	"github.com/RaymondCode/simple-demo/model"
)

func GetPublishList(id int64) []model.Video {
	db := connection.GetDB()
	var videos []model.Video
	db.Where("author_id = ?",id).Find(&videos)
	return videos
}
