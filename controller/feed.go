package controller

import (
	"github.com/RaymondCode/simple-demo/dao/redis"
	"github.com/RaymondCode/simple-demo/model"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type FeedResponse struct {
	model.Response
	VideoList []model.Video `json:"video_list,omitempty"`
	NextTime  int64   `json:"next_time,omitempty"`
}


// Feed same demo video list for every request
func Feed(c *gin.Context) {
	token := c.Query("token")
	var user model.User
	if token != "" {
		user = redis.GetFromRedis(token)
		//usersLoginInfo[token] = user
	}
	if user.Id != 0 {
		c.JSON(http.StatusOK, FeedResponse{
			Response:  model.Response{StatusCode: 0},
			VideoList: service.SelectAllVideoByFavorite(user.Id),
			//VideoList: DemoVideos,
			NextTime:  time.Now().Unix(),
		})
	}else {
		c.JSON(http.StatusOK, FeedResponse{
			Response:  model.Response{StatusCode: 0},
			VideoList: service.SelectAllVideo(),
			//VideoList: DemoVideos,
			NextTime:  time.Now().Unix(),
		})
	}

}
