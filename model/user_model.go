package model

import (
	"Gin_book/database"
)

//注册用户 结构
type User struct {
	Username   string `json:username`
	Password   string `json:password`
	Repassword string `json:repassword`
	Phone      string `json:phone`
}

//--------------数据库操作-----------------

//插入
func InsertUser(user User) (int64, error) {
	return database.ModifyDB("insert into users(username,password,repassword,phone) values (?,?,?,?)",
		user.Username, user.Password, user.Repassword, user.Phone)
}

//按条件查询
//func QueryUserWightCon(con string) int {
//	sql := fmt.Sprintf("select id from users %s", con)
//	fmt.Println(sql)
//	row := database.QueryRowDB(sql)
//	id := 0
//	row.Scan(&id)
//	return id
//}

////根据用户名查询id
//func QueryUserWithUsername(username string) int {
//	sql := fmt.Sprintf("where username='%s'", username)
//	return QueryUserWightCon(sql)
//}

////根据用户名和密码，查询id
//func QueryUserWithParam(username ,password string)int{
//	sql:=fmt.Sprintf("where username='%s' and password='%s'",username,password)
//	return QueryUserWightCon(sql)
//}
