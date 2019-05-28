package controller

import (
	m "Gin_blog/model"
	"Gin_blog/util"
	"fmt"
	g "github.com/gin-gonic/gin"
	"net/http"
)

func RegisterGet(c *g.Context) {
	//返回html
	c.HTML(http.StatusOK, "index.html", g.H{"title": "注册页"})
}
func RegisterPost(c *g.Context) {
	var User m.User
	if err := c.ShouldBindJSON(&User); err != nil {
		//错误返回
		c.JSON(http.StatusBadRequest, g.H{"error": err.Error()})
		return
	}
	//
	fmt.Printf("user:1 %s %s %s %s", User.Username, User.Password, User.Phone, User.Repassword)
	//
	//c.JSON(http.StatusOK, g.H{"starts": "azhen "})
	////注册之前先判断该用户名是否已经被注册，如果已经注册，返回错误
	//id := m.QueryUserWithUsername(username)
	//fmt.Println("id:", id)
	//if id > 0 {
	//	c.JSON(http.StatusOK, g.H{"code": 0, "message": "用户名已经存在"})
	//	return
	//}

	//注册用户名和密码
	//存储的密码是md5后的数据，那么在登录的验证的时候，也是需要将用户的密码md5之后和数据库里面的密码进行判断
	User.Password = util.MD5(User.Password)
	fmt.Println("md5后：", User.Password)
	//
	user := m.User{
		Username:   User.Username,
		Password:   User.Password,
		Repassword: User.Repassword,
		Phone:      User.Phone,
	}
	_, err := m.InsertUser(user)
	if err != nil {
		c.JSON(http.StatusOK, g.H{"code": 0, "message": "注册失败"})
	} else {
		c.JSON(http.StatusOK, g.H{"code": 1, "message": "注册成功"})
	}
}
