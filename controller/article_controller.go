package controller

import (
	m "Gin_book/model"
	"fmt"
	g "github.com/gin-gonic/gin"
	"net/http"
)

/*
当访问/add路径的时候回触发AddArticleGet方法
响应的页面是通过HTML
*/
func AddArticleGet(c *g.Context) {
	//这里是具体实现
	//获取session
	//islog := GetSession(c)

	c.HTML(http.StatusOK, "write_article.html", g.H{"IsLogin": "islogin"})
}

func AddArticlePost(c *g.Context) {
	var Art m.Article
	if err := c.ShouldBindJSON(&Art); err != nil {
		//错误返回
		c.JSON(http.StatusBadRequest, g.H{"error": err.Error()})
		return
	}
	fmt.Printf("test id %s", Art.Short)
	//实例化model，将它出入到数据库中
	art := m.Article{
		Art.Id,
		Art.Title,
		Art.Tags,
		Art.Short,
		Art.Content,
		Art.Author,
		Art.Createtime,
		Art.Status,
	}
	_, err := m.AddArticle(art)

	//返回数据给浏览器
	response := g.H{}
	if err == nil {
		//无误
		response = g.H{"code": 1, "message": "ok"}
	} else {
		response = g.H{"code": 0, "message": "error"}
	}

	c.JSON(http.StatusOK, response)

}
