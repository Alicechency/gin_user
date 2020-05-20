package models

import (
	"fmt"
	db "gin_login/database"
	"gin_login/utils"
)


type User struct {
	Id        int
	UserName string
	PassWord  string
	Email string
	CreateTime int64
}

func (u * User) AddUser() (id int64, err error) {
	return utils.ModifyDB("INSERT INTO users(id, username, password, email,createtime) VALUES (?, ?, ?, ?, ?)", u.Id, u.UserName, u.PassWord, u.Email,u.CreateTime)

}

//按条件查询
func QueryUserWightCon(con string) int {
	sql := fmt.Sprintf("select id from users %s", con)
	row := db.QueryRowDB(sql)
	id := 0
	row.Scan(&id)
	return id
}

//根据用户名查询id
func QueryUserWithUsername(username string) int {
	sql := fmt.Sprintf("where username='%s'", username)
	return QueryUserWightCon(sql)
}

//根据用户名和密码，查询id
func QueryUserWithParam(username ,password string)int{
	sql:=fmt.Sprintf("where username='%s' and password='%s'",username,password)
	return QueryUserWightCon(sql)
}

//修改数据
func UpdateUser(user User) (int64, error) {
	//数据库操作
	return utils.ModifyDB("update users set password=? where id=?",
		user.PassWord,  user.Id)
}


