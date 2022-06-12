package connection

import (
	"github.com/RaymondCode/simple-demo/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//定义全局的db对象，我们执行数据库操作主要通过他实现。
var _db *gorm.DB

//包初始化函数，golang特性，每个包初始化的时候会自动执行init函数，这里用来初始化gorm。
func init() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/douyin?charset=utf8mb4&parseTime=True&loc=Local"
	// 声明err变量，下面不能使用:=赋值运算符，否则_db变量会当成局部变量，导致外部无法访问_db变量
	var err error
	//连接MYSQL, 获得DB类型实例，用于后面的数据库读写操作。
	_db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	_db.AutoMigrate(model.User{})
	_db.AutoMigrate(&model.Video{})
	_db.AutoMigrate(&model.LoginInfo{})
	_db.AutoMigrate(&model.Follow{})
	_db.AutoMigrate(&model.Comment{})
	_db.AutoMigrate(&model.FavoriteVideo{})
	_db.AutoMigrate(&model.Token{})
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}
	sqlDB, _ := _db.DB()
	sqlDB.SetMaxOpenConns(100)  //设置数据库连接池最大连接数
	sqlDB.SetMaxIdleConns(20) //连接池最大允许的空闲连接数，如果没有sql任务需要执行的连接数大于20，超过的连接会被连接池关闭。
}

func GetDB() *gorm.DB {
	return _db
}