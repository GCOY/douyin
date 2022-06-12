package service

import (
	"github.com/RaymondCode/simple-demo/dao"
	"github.com/RaymondCode/simple-demo/model"
)

func FollowAction(followID int64,followerID int64)  {
	dao.FollowAction(followID,followerID)
}

func CancelFollow(followID int64,followerID int64)  {
	dao.CancelFollow(followID,followerID)
}

func GetFollowList(followerID int64) []model.User {
	users := dao.GetFollow(followerID)
	for i:=0;i< len(users);i++ {
		users[i].IsFollow = true
	}
	return users
}

func GetFollowerList(followID int64) []model.User {
	users := dao.GetFollower(followID)
	followUsers := GetFollowList(followID)
	for i := 0;i< len(followUsers);i++ {
		for j := 0 ;j< len(users);j++ {
			if followUsers[i].Id == users[j].Id {
				users[j].IsFollow = true
			}
		}
	}
	return users
}
