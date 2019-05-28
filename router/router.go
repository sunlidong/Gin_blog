package router

import (
	cf "Gin_book/conf"
	ctl "Gin_book/controller"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() *gin.Engine {

	router := gin.Default()
	router.LoadHTMLGlob(cf.HtmlFile)
	router.Static("/static", cf.StaticFile)
	{
		//路由組
		api := router.Group("api")
		{
			//注册：
			api.GET("/register", ctl.RegisterGet)
			api.POST("/register", ctl.RegisterPost)
			//登录
			api.GET("/login", ctl.LoginGet)
			api.POST("/login", ctl.LoginGet)

			//首页
			api.GET("/", ctl.LoginGet)

			//退出
			api.GET("/exit", ctl.LoginGet)
			//
		}

		//  文章类
		v1 := api.Group("/v1")
		{
			v1.GET("getuser", func(c *gin.Context) {
				//
				fmt.Println("test1")
				c.JSON(http.StatusOK, gin.H{"starts": "測試成功"})
			})
			//
			//写文章
			v1.GET("/add", ctl.AddArticleGet)
			v1.POST("/add", ctl.AddArticlePost)
			//显示文章内容
			//v1.GET("/show/:id", ctl.ShowArticleGet)
			//
			////更新文章
			//v1.GET("/update",ctl.UpdateArticleGet)
			//v1.POST("/update",ctl.UpdateArticlePost)
			//
			//
			////删除文章
			//v1.GET("/delete",controllers.DeleteArticleGet)
		}
	}
	return router
}
