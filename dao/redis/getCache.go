package redis

import (
	"encoding/json"
	"fmt"
	"github.com/RaymondCode/simple-demo/connection"
	"github.com/RaymondCode/simple-demo/dao"
	"github.com/RaymondCode/simple-demo/model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gomodule/redigo/redis"
)


func GetFromRedis(token string) model.User {
	rdb := connection.RedisInit()
	defer connection.RedisClose()
	var user model.User
	r, err := redis.Bytes(rdb.Do("get", token))
	if len(r) > 0 {
		err = json.Unmarshal(r, &user)
		if err != nil {
			fmt.Println(err)
		}
		return user
	}else {
		userInMysql := dao.SelectUserByToken(token)
		if userInMysql.Id != 0 {
			data, err := json.Marshal(userInMysql)
			if err != nil {
				fmt.Println(err)
			}
			SetToRedis(token,data)
			return userInMysql
		}
	}
	return user
}

func SetToRedis(token string,data []byte)  {
	rdb := connection.RedisInit()
	defer connection.RedisClose()
	var err error
	_, err = rdb.Do("set", token, data,"EX",3600)
	if err != nil {
		fmt.Println(err)
		return
	}
}

