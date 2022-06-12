package service

import (
	"github.com/RaymondCode/simple-demo/dao"
	"github.com/RaymondCode/simple-demo/model"
)

func CancelFavorite(userID int64, videoID int64)  {
	dao.CancelFavorite(userID,videoID)
}

func IsFavorite(userID int64, videoID int64)  {
	dao.IsFavorite(userID,videoID)
}

func FavoriteList(userID int64) []model.Video {
	return dao.FavoriteList(userID)
}
