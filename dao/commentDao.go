package dao

import (
	"github.com/RaymondCode/simple-demo/connection"
	"github.com/RaymondCode/simple-demo/model"
	"sync"
)

func CreateComment(userID int64,context string,createDate string,videoID int64)  {
	var mutex sync.Mutex
	db := connection.GetDB()
	var comment model.Comment
	comment.UserID = userID
	comment.Content = context
	comment.CreateDate = createDate
	comment.VideoID = videoID
	var video model.Video
	db.Where("id = ?",videoID).Find(&video)
	mutex.Lock()
	tx := db.Begin()
	if err := tx.Create(&comment).Error;err != nil {
		tx.Rollback()
		return
	}
	if err := tx.Model(&video).Where("id = ?",videoID).Update("comment_count",video.CommentCount+1).Error;err != nil {
		tx.Rollback()
		return
	}
	tx.Commit()
	mutex.Unlock()
}

func CancelComment(commentID int64,videoID int64)  {
	var mutex sync.Mutex
	db := connection.GetDB()
	var comment model.Comment
	var video model.Video
	db.Where("id = ?",videoID).Find(&video)
	mutex.Lock()
	tx := db.Begin()
	if err := tx.Where("id = ?",commentID).Delete(&comment).Error;err != nil {
		tx.Rollback()
		return
	}
	if err := tx.Model(&video).Where("id = ?",videoID).Update("comment_count",video.CommentCount-1).Error;err != nil {
		tx.Rollback()
		return
	}
	tx.Commit()
	mutex.Unlock()
}

func GetCommentList(videoID int64) []model.Comment {
	db := connection.GetDB()
	var commentList []model.Comment
	db.Where("video_id = ?",videoID).Preload("User").Find(&commentList)
	return commentList
}
