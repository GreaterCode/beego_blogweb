package utils

import (
	"database/sql"
	"fmt"
	"log"
)

var db *sql.DB

func InitMysql() {
	fmt.Println("Init Mysql ......")
	if db == nil {
		db, _ = sql.Open("mysql", "root@")
		CreateTableWithUser()
		CreateTableWithArticle()
		CreateTableWithAlbum()
	}
}

func CreateTableWithAlbum() {
	sql := `create table if not exists album(
		id int(4) primary key auto_increment not null,
		filepath varchar(255),
		filename varchar(64),
		status int(4),
		createtime int(10)
		);`
	ModifyDB(sql)
}

// 创建文章表
func CreateTableWithArticle() {
	sql := `create table if not exists article(
		id int(4) primary key auto_increment not null,
		title varchar(30),
		author varchar(20),
		tags varchar(30),
		short varchar(255),
		content longtext,
		createtime int(10)
		);`
	ModifyDB(sql)
}

// 创建用户表
func CreateTableWithUser() {
	sql := `create table if not exists user(
		id  int(4) primary key auto_increment not null,
		username varchar(64),
		password varchar(64),
		status int(4),
		createtime int(10)
	)`
	ModifyDB(sql)
}

// 操作数据库
func ModifyDB(sql string, args ...interface{}) (int64, error) {
	result, err := db.Exec(sql, args...)
	if err != nil {
		log.Println(err)
	}
	count, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
		return 0, err
	}

	return count, nil
}

// 查询数据库
func QueryRowDB(sql string) *sql.Row {
	return db.QueryRow(sql)
}
