package db

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"time"
	"z-blog-go/configs"
)

var DB *sql.DB

func Open() (*sql.DB, error) {
	app := configs.GetApp()
	config := app.Database
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable fallback_application_name=z-blog-go", config.Host, config.Port, config.User, config.Password, config.Name)
	db, err := sql.Open("postgres", psqlInfo)

	//connStr := "postgres://testuser:testpwd@127.0.0.1/testdb?sslmode=verify-full"
	//db, err := sql.Open("postgres", connStr)

	if err != nil {
		return nil, err
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	//连接池中“打开”连接(使用中+空闲连接)数量的上限，默认是无限的
	//PostgreSQL 默认最多100个打开连接的限制，达到限制驱动返回"sorry, too many clients already"错误
	//PostgreSQL 最大打开连接数限制可以在 postgresql.conf 文件中 max_connections 设置来更改
	db.SetMaxOpenConns(20)
	//连接池中空闲连接数的上限，默认是 2
	//MySQL 会自动关闭任何 8 小时未使用的连接
	db.SetMaxIdleConns(5)
	//一个连接保持可用的最长时间，默认没有限制，永久可用
	db.SetConnMaxLifetime(8 * time.Hour)
	//1.15版本引入，自上次使用以后在连接池中空闲了 1 小时的任何连接都将被标记为过期并被后台清理操作删除
	db.SetConnMaxIdleTime(1 * time.Hour)

	//创建一个具有 5 秒超时期限的上下文。
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}
	DB = db
	return db, nil
}
