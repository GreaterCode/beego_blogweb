package models

import (
	"beego_blogweb/utils"
	"fmt"
)

type User struct {
	Id         int
	Username   string
	Password   string
	Status     int // 0 正常状态， 1删除
	Createtime int64
}

// 插入
func InsertUser(user User) (int64, error) {
	return utils.ModifyDB("insert into users(username, password, status, createtime) values(?,?,??)",
		user.Username, user.Password, user.Status, user.Createtime)
}

// 按照条件查询
func QueryUserWithCon(con string) int {
	sql := fmt.Sprintf("select id from user %s", con)
	fmt.Println(sql)
	row := utils.QueryRowDB(sql)
	id := 0
	row.Scan(&id)

	return id
}

// 根据用户名查询id
func QueryUserWithUsername(username string) int {
	sql := fmt.Sprintf("where username =%s", username)
	return QueryUserWithCon(sql)
}

// 根据用户名和密码查询
func QueryUserWithParam(username, password string) int {
	sql := fmt.Sprintf("where username='%s' and password=%s", username, password)
	return QueryUserWithCon(sql)

}
