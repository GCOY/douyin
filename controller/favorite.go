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

// FavoriteAction no practical effect, just check if token is valid
func FavoriteAction(c *gin.Context) {
	token := c.Query("token")
	var user model.User
	if token != "" {
		user = redis.GetFromRedis(token)
		//usersLoginInfo[token] = user
	}
	actionType := c.Query("action_type")
	videoId,err := strconv.ParseInt(c.Query("video_id"),10,64)
	if err!=nil{
		fmt.Println("类型转换失败")
	}
	if user.Id != 0 {
		if actionType == "1"{
			service.IsFavorite(user.Id,videoId)
		}
		if actionType == "2"{
			service.CancelFavorite(user.Id,videoId)
		}
		c.JSON(http.StatusOK, model.Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, model.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

// FavoriteList all users have same favorite video list
func FavoriteList(c *gin.Context) {
	token := c.Query("token")
	userid,err := strconv.ParseInt(c.Query("user_id"),10,64)
	var user model.User
	if token != "" {
		user = redis.GetFromRedis(token)
		//usersLoginInfo[token] = user
	}
	if err!=nil{
		fmt.Println("类型转换失败")
	}
	if userid != user.Id {
		c.JSON(http.StatusOK, VideoListResponse{
			Response: model.Response{
				StatusCode: 0,
			},
			VideoList: service.FavoriteList(userid),
			//VideoList: DemoVideos,
		})
	} else {
		c.JSON(http.StatusOK, VideoListResponse{
			Response: model.Response{
				StatusCode: 0,
			},
			VideoList: service.FavoriteList(user.Id),
			//VideoList: DemoVideos,
		})
	}
}
