package controller

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/dao/redis"
	"github.com/RaymondCode/simple-demo/model"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserListResponse struct {
	model.Response
	UserList []model.User `json:"user_list"`
}

// RelationAction no practical effect, just check if token is valid
func RelationAction(c *gin.Context) {
	token := c.Query("token")
	var user model.User
	if token != "" {
		user = redis.GetFromRedis(token)
		//usersLoginInfo[token] = user
	}
	actionType := c.Query("action_type")
	toUserId,err := strconv.ParseInt(c.Query("to_user_id"),10,64)
	if err!=nil{
		fmt.Println("类型转换失败")
	}
	if user.Id != 0 {
		if actionType == "1" {
			service.FollowAction(toUserId,user.Id)
		}
		if actionType == "2" {
			service.CancelFollow(toUserId,user.Id)
		}
		c.JSON(http.StatusOK, model.Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, model.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

// FollowList all users have same follow list
func FollowList(c *gin.Context) {
	token := c.Query("token")
	var user model.User
	if token != "" {
		user = redis.GetFromRedis(token)
		//usersLoginInfo[token] = user
	}
	userid,err := strconv.ParseInt(c.Query("user_id"),10,64)
	if err!=nil{
		fmt.Println("类型转换失败")
	}
	if userid != user.Id {
		c.JSON(http.StatusOK, UserListResponse{
			Response: model.Response{
				StatusCode: 0,
			},
			UserList: service.GetFollowList(userid),
			//UserList: []model.User{DemoUser},
		})
	}else {
		c.JSON(http.StatusOK, UserListResponse{
			Response: model.Response{
				StatusCode: 0,
			},
			UserList: service.GetFollowList(user.Id),
			//UserList: []model.User{DemoUser},
		})
	}
}

// FollowerList all users have same follower list
func FollowerList(c *gin.Context) {
	token := c.Query("token")
	var user model.User
	if token != "" {
		user = redis.GetFromRedis(token)
		//usersLoginInfo[token] = user
	}
	userid,err := strconv.ParseInt(c.Query("user_id"),10,64)
	if err!=nil{
		fmt.Println("类型转换失败")
	}
	if userid != user.Id {
		c.JSON(http.StatusOK, UserListResponse{
			Response: model.Response{
				StatusCode: 0,
			},
			UserList: service.GetFollowerList(userid),
			//UserList: []model.User{DemoUser},
		})
	}else {
		c.JSON(http.StatusOK, UserListResponse{
			Response: model.Response{
				StatusCode: 0,
			},
			UserList: service.GetFollowerList(user.Id),
			//UserList: []model.User{DemoUser},
		})
	}
}
