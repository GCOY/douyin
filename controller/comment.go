package controller

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/dao/redis"
	"github.com/RaymondCode/simple-demo/model"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type CommentListResponse struct {
	model.Response
	CommentList []model.Comment `json:"comment_list,omitempty"`
}

type CommentActionResponse struct {
	model.Response
	Comment model.Comment `json:"comment,omitempty"`
}

// CommentAction no practical effect, just check if token is valid
func CommentAction(c *gin.Context) {
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
		if actionType == "1" {
			text := c.Query("comment_text")
			timeNow := time.Now().Format("01-02")
			//nowTime := month + "-" +day
			service.CreateComment(user.Id,text,timeNow,videoId)
			c.JSON(http.StatusOK, CommentActionResponse{Response: model.Response{StatusCode: 0},
				Comment: model.Comment{
					Id:         1,
					User:       user,
					Content:    text,
					CreateDate: timeNow,
				}})
			return
		}
		if actionType == "2" {
			commentId := c.Query("comment_id")
			commentID,err := strconv.ParseInt(commentId,10,64)
			if err!=nil{
				fmt.Println("类型转换失败")
			}
			service.CancelComment(commentID,videoId)
		}
		c.JSON(http.StatusOK, model.Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, model.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

// CommentList all videos have same demo comment list
func CommentList(c *gin.Context) {
	videoId,err := strconv.ParseInt(c.Query("video_id"),10,64)
	if err!=nil{
		fmt.Println("类型转换失败")
	}
	c.JSON(http.StatusOK, CommentListResponse{
		Response:    model.Response{StatusCode: 0},
		CommentList: service.GetCommentList(videoId),
		//CommentList: DemoComments,
	})
}
