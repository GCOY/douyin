package dao

import (
	"github.com/RaymondCode/simple-demo/connection"
	"github.com/RaymondCode/simple-demo/model"
	"sync"
)

func FollowAction(followID int64,followerID int64)  {
	var mutex sync.Mutex
	db := connection.GetDB()
	var follow model.Follow
	//FollowerUser:用户
	//FollowUser:对方用户
	var FollowUser model.User
	var FollowerUser model.User
	db.Where("id = ?",followID).Find(&FollowUser)
	db.Where("id = ?",followerID).Find(&FollowerUser)
	//向follows表中查信息，没有就创建
	mutex.Lock()
	tx := db.Begin()
	if err := tx.Where("follow_id = ? AND follower_id = ?",followID,followerID).First(&follow).Error;err != nil{
		//创建follow记录，关注状态设为true
		if err := tx.Create(&model.Follow{FollowID: followID, FollowerID: followerID, FollowStatus: true}).Error; err != nil {
			tx.Rollback()
			return
		}
		//将用户的关注数+1，将对方用户的粉丝数+1
		if err := tx.Model(&FollowUser).Where("id = ?",followID).Update("follower_count",FollowUser.FollowerCount+1).Error; err != nil {
			tx.Rollback()
			return
		}
		if err := tx.Model(&FollowerUser).Where("id = ?",followerID).Update("follow_count",FollowerUser.FollowCount+1).Error; err != nil {
			tx.Rollback()
			return
		}
	}else {
		//处理的是记录存在的情况
		//关注状态设为true
		if err := tx.Model(&follow).Where("follow_id = ? AND follower_id = ?",followID,followerID).Update("follow_status",true).Error; err != nil {
			tx.Rollback()
			return
		}
		if err := tx.Model(&FollowUser).Where("id = ?",followID).Update("follower_count",FollowUser.FollowerCount+1).Error; err != nil {
			tx.Rollback()
			return
		}
		if err := tx.Model(&FollowerUser).Where("id = ?",followerID).Update("follow_count",FollowerUser.FollowCount+1).Error; err != nil {
			tx.Rollback()
			return
		}
	}
	tx.Commit()
	mutex.Unlock()
}

func CancelFollow(followID int64,followerID int64)  {
	var mutex sync.Mutex
	db := connection.GetDB()
	var follow model.Follow
	var FollowUser model.User
	var FollowerUser model.User
	//follower:本地用户
	//follow:对方用户
	db.Where("id = ?",followID).Find(&FollowUser)
	db.Where("id = ?",followerID).Find(&FollowerUser)
	mutex.Lock()
	tx := db.Begin()
	if err:=tx.Model(&follow).Where("follow_id = ? AND follower_id = ?",followID,followerID).Update("follow_status",false).Error;err != nil {
		tx.Rollback()
		return
	}
	if err:=tx.Model(&FollowUser).Where("id = ?",followID).Update("follower_count",FollowUser.FollowerCount-1).Error;err != nil {
		tx.Rollback()
		return
	}
	if err:=tx.Model(&FollowerUser).Where("id = ?",followerID).Update("follow_count",FollowerUser.FollowCount-1).Error;err != nil {
		tx.Rollback()
		return
	}
	tx.Commit()
	mutex.Unlock()
}

func GetFollow(followerID int64) []model.User {
	db := connection.GetDB()
	var users []model.User
	var followID []int64
	//db.Table("favorite_videos").Select("video_id").Where("user_id = ? AND favorite_status = ?",userID,1).Scan(&videoID)
	db.Table("follows").Select("follow_id").Where("follower_id = ? AND follow_status = ?",followerID,true).Scan(&followID)
	//db.Where("id IN ?",videoID).Preload("Author").Find(&videos)
	db.Where("id IN ?",followID).Find(&users)
	return users
}

func GetFollower(followID int64) []model.User {
	db := connection.GetDB()
	var users []model.User
	var followerID []int64
	db.Table("follows").Select("follower_id").Where("follow_id = ? AND follow_status = ?",followID,1).Scan(&followerID)
	db.Where("id IN ?",followerID).Find(&users)
	return users
}

