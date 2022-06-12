package controller

import (
	"github.com/RaymondCode/simple-demo/dao"
	"github.com/RaymondCode/simple-demo/dao/redis"
	"github.com/RaymondCode/simple-demo/middleware"
	"github.com/RaymondCode/simple-demo/model"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

//var usersLoginInfo = make(map[string] model.User)

type UserLoginResponse struct {
	model.Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	model.Response
	User model.User `json:"user"`
}

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	//usersLoginInfo = service.GetUserMap(username)
	//token := username + password

	if exist := service.JudgeUser(username);!exist {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: model.Response{StatusCode: 1, StatusMsg: "User already exist"},
		})
	}else {
		//atomic.AddInt64(&userIdSequence, 1)
		service.UserInfoRegister(username,password)
		userId := service.GetUserID(username)
		service.UserRegister(username,userId)
		service.SaveToken(username,password,userId)
		//newUser := model.User{
		//	Id:   userId,
		//	Name: username,
		//}
		////usersLoginInfo = make(map[string]model.User)
		//usersLoginInfo[token] = newUser
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: model.Response{StatusCode: 0},
			UserId:   userId,
			Token:    middleware.MD5(username + password),
		})
	}
	//if _, exist := usersLoginInfo[token]; exist {
	//	c.JSON(http.StatusOK, UserLoginResponse{
	//		Response: model.Response{StatusCode: 1, StatusMsg: "User already exist"},
	//	})
	//}
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	//usersLoginInfo = service.GetUserMap(username)
	user := dao.SelectUserByName(username)
	token := middleware.MD5(username + password)

	if user.Id != 0 {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: model.Response{StatusCode: 0},
			UserId:   user.Id,
			Token:    token,
		})
	} else {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: model.Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}

func UserInfo(c *gin.Context) {
	token := c.Query("token")
	var user model.User
	if token != "" {
		user = redis.GetFromRedis(token)
		//usersLoginInfo[token] = user
	}
	if user.Id != 0 {
		c.JSON(http.StatusOK, UserResponse{
			Response: model.Response{StatusCode: 0},
			User:     user,
		})
	} else {
		c.JSON(http.StatusOK, UserResponse{
			Response: model.Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}
