package service

import (
	"github.com/RaymondCode/simple-demo/dao"
	"github.com/RaymondCode/simple-demo/model"
)

func CreateComment(userID int64,context string,createDate string,videoID int64)  {
	dao.CreateComment(userID,context,createDate,videoID)
}

func CancelComment(commentID int64,videoID int64)  {
	dao.CancelComment(commentID,videoID)
}

func GetCommentList(videoID int64) []model.Comment {
	return dao.GetCommentList(videoID)
}
