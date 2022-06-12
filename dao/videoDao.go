package dao

import (
	"github.com/RaymondCode/simple-demo/connection"
	"github.com/RaymondCode/simple-demo/model"
	"sync"
	"time"
)


//根据用户id查询videos表视频
func SelectAllVideoByID(userID int64) []model.FavoriteVideo {
	db := connection.GetDB()
	var favorites []model.FavoriteVideo
	db.Where("user_id = ?" ,userID).Find(&favorites)
	return favorites
}

//发布时间倒叙查询videos表视频
func SelectAllVideoByPublishTimeDesc() []model.Video{
	timeNow := time.Now().Unix()
	db := connection.GetDB()
	var videos []model.Video
	db.Where("publish_time < ?",timeNow).Limit(30).Order("publish_time desc").Preload("Author").Find(&videos)
	return videos
}

//向videos表保存视频
func SaveVideo(userid int64, filename string ,cover string,nowTime int64) {
	var mutex sync.Mutex
	db := connection.GetDB()
	//获取url前缀
	vnamePrefix := "https://my-douyin.oss-cn-hangzhou.aliyuncs.com/douyinvideo/"
	pnamePrefix := "https://my-douyin.oss-cn-hangzhou.aliyuncs.com/douyincover/"
	//namePrefix := "http://192.168.11.32:8080/static/"
	playUrl := vnamePrefix + filename
	coverUrl := pnamePrefix + cover
	mutex.Lock()
	tx := db.Begin()
	video := model.Video{AuthorID: userid, PlayUrl: playUrl,CoverUrl: coverUrl,PublishTime: nowTime}
	if err := tx.Create(&video).Error; err != nil {
		tx.Rollback()
		return
	}
	tx.Commit()
	mutex.Unlock()
}