package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
)

var (
	Db *sql.DB
	R  *redis.Client
)

func init() {
	username := "root"       // 账号
	password := "1051140642" // 密码
	host := "127.0.0.1"      // 数据库地址，可以是Ip或者域名
	port := 3306             // 数据库端口
	dbName := "db_task"      // 数据库名

	// MYSQL dsn格式： {username}:{password}@tcp({host}:{port})/{dbName}?charset=utf8&parseTime=True&loc=Local
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=True&loc=Local&charset=utf8mb4&collation=utf8mb4_unicode_ci",
		username, password, host, port, dbName)

	// 连接MYSQL
	db, err := sql.Open("mysql", dsn)
	db.SetMaxIdleConns(100)
	db.SetMaxOpenConns(120)
	db.SetConnMaxLifetime(2 * time.Second)
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	Db = db

	R = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // redis地址
		Password: "",               // redis密码，没有则留空
		DB:       0,                // 默认数据库，默认是0
	})

	// R.Set("test", time.Now().String(), 0)
}
