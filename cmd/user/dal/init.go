package dal

import (
	"database/sql"
	"simple-douyin/cmd/user/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormopentracing "gorm.io/plugin/opentracing"
)

var DB *gorm.DB

func GetDB() *gorm.DB {
	return DB
}

const (
	MYSQLDSN = "root:Zm1021ok@tcp(127.0.0.1:3306)/"
	DBName = "douyinDB"
)

// 初始化数据库创建函数
func initializeMysqlDB() {
	// 连接MySQL的服务器
	db, err := sql.Open("mysql", MYSQLDSN)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	// 构建SQL语句
	sqlStr := "CREATE DATABASE IF NOT EXISTS "+ DBName + " Character Set utf8mb4"
	// 创建数据库
	if _, err = db.Exec(sqlStr); err != nil {
		panic("数据库创建失败！")
	}
}

// Init init DB
func init() {
	initializeMysqlDB()
	var err error
	DB, err = gorm.Open(mysql.Open(MYSQLDSN+DBName),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}

	if err = DB.Use(gormopentracing.New()); err != nil {
		panic(err)
	}

	DB.AutoMigrate(new(model.User), new(model.Video), new(model.Comment))
}
