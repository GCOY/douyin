package controller

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/dao/aliyunOSS"
	"github.com/RaymondCode/simple-demo/dao/redis"
	"github.com/RaymondCode/simple-demo/middleware"
	"github.com/RaymondCode/simple-demo/model"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
	"strconv"
	"time"
)

type VideoListResponse struct {
	model.Response
	VideoList []model.Video `json:"video_list"`
}

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
	token := c.PostForm("token")
	var user model.User
	if token != "" {
		user = redis.GetFromRedis(token)
		//usersLoginInfo[token] = user
	}
	if user.Id == 0 {
		c.JSON(http.StatusOK, model.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}

	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	nowTime := time.Now().Unix()
	timeNow := strconv.FormatInt(nowTime, 10)
	filename := filepath.Base(data.Filename)
	middlefilename := timeNow + filename
	//user := usersLoginInfo[token]
	finalName := fmt.Sprintf("%s", middlefilename)
	saveFile := filepath.Join("./public/", finalName)
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	pictureName := finalName[0:len(finalName)-3]
	cover := pictureName + "jpeg"
	//获取url前缀
	fmt.Println(cover)

	service.SaveVideo(user.Id, finalName,cover, nowTime)
	c.JSON(http.StatusOK, model.Response{
		StatusCode: 0,
		StatusMsg:  finalName + " uploaded successfully",
	})
	middleware.SavePicture(finalName,cover)
	aliyunOSS.UploadVideoFile(finalName)
	aliyunOSS.UploadFileFile(cover)
}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
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
			VideoList: service.GetPublishList(userid),
			//VideoList: DemoVideos,
		})
	} else {
		c.JSON(http.StatusOK, VideoListResponse{
			Response: model.Response{
				StatusCode: 0,
			},
			VideoList: service.GetPublishList(user.Id),
			//VideoList: DemoVideos,
		})
	}
}
