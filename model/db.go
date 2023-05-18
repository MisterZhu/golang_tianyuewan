package model

import (
	"fmt"
	"gindiary/util"

	"gorm.io/driver/mysql"
	"gorm.io/gorm/schema"

	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func InitDb() {
	//配置MySQL连接参数
	username := util.DBUsername //账号
	password := util.DBPassword //密码
	host := util.DBHost         //数据库地址，可以是Ip或者域名
	port := util.DBPort         //数据库端口
	Dbname := util.DBName       //数据库名
	charset := util.DBCharset   //编码格式
	loc := util.DBLoc           //本地
	timeout := util.DBTimeout   //连接超时，10秒

	//拼接下dsn参数, dsn格式可以参考上面的语法，这里使用Sprintf动态拼接dsn参数，因为一般数据库连接参数，我们都是保存在配置文件里面，需要从配置文件加载参数，然后拼接dsn。
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&loc=%s&timeout=%s&parseTime=true", username, password, host, port, Dbname, charset, loc, timeout)
	//连接MYSQL, 获得DB类型实例，用于后面的数据库读写操作。
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic("failed to connect database, err: " + err.Error())
	}
	db.AutoMigrate(&User{}, &Article{}, &Category{}, &XcxUser{}, &XcxAnalyModel{})

}
