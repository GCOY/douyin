package dao

import (
	"github.com/RaymondCode/simple-demo/connection"
	"github.com/RaymondCode/simple-demo/model"
	"sync"
)

func CancelFavorite(userID int64, videoID int64)  {
	var mutex sync.Mutex
	db := connection.GetDB()
	var favorite model.FavoriteVideo
	var video model.Video
	db.Where("id = ?",videoID).Find(&video)
	mutex.Lock()
	tx := db.Begin()
	//喜欢状态设为false
	if err := tx.Model(&favorite).Where("user_id = ? AND video_id = ?",userID,videoID).Update("favorite_status",false).Error;err != nil {
		tx.Rollback()
		return
	}
	//对应video的favorite_count-1
	if err := tx.Model(&video).Where("id = ?",videoID).Update("favorite_count",video.FavoriteCount-1).Error;err != nil {
		tx.Rollback()
		return
	}
	tx.Commit()
	mutex.Unlock()
}

func IsFavorite(userID int64, videoID int64)  {
	var mutex sync.Mutex
	db := connection.GetDB()
	var favorite model.FavoriteVideo
	var video model.Video
	db.Where("id = ?",videoID).Find(&video)
	mutex.Lock()
	tx := db.Begin()
	//从favorites表获取记录，没有就创建
	if err := tx.Where("user_id = ? AND video_id = ?",userID,videoID).First(&favorite).Error;err != nil{
		//喜欢状态设为true
		if err := tx.Create(&model.FavoriteVideo{UserID: userID,VideoID: videoID, FavoriteStatus: true}).Error;err != nil {
			tx.Rollback()
			return
		}
		//对应video的favorite_count+1
		if err := tx.Model(&video).Where("id = ?",videoID).Update("favorite_count",video.FavoriteCount+1).Error;err != nil {
			tx.Rollback()
			return
		}
	}else {
		//喜欢状态设为true
		if err := tx.Model(&favorite).Where("user_id = ? AND video_id = ?",userID,videoID).Update("favorite_status",true).Error;err != nil {
			tx.Rollback()
			return
		}
		//对应video的favorite_count+1
		if err := tx.Model(&video).Where("id = ?",videoID).Update("favorite_count",video.FavoriteCount+1).Error;err != nil {
			tx.Rollback()
			return
		}
	}
	tx.Commit()
	mutex.Unlock()
}

func FavoriteList(userID int64) []model.Video {
	db := connection.GetDB()
	var videoID []int64
	//videoID := db.Select("video_id").Where("user_id = ?",userID).Table("favorite_videos")
	db.Table("favorite_videos").Select("video_id").Where("user_id = ? AND favorite_status = ?",userID,1).Scan(&videoID)
	//db.Model(&favorite).Where("user_id = ?",userID).Pluck("video_id",videoID)
	var videos []model.Video
	db.Where("id IN ?",videoID).Preload("Author").Find(&videos)
	return videos
}