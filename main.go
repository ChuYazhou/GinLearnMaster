package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

type Article struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func main() {

	// 路由初始化
	r := gin.Default()

	r.LoadHTMLGlob("templates/*")

	// 路由注册
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	r.GET("/ping", func(c *gin.Context) {
		c.String(201, "我是hello world界面")
	})

	r.GET("/news", func(c *gin.Context) {
		c.String(http.StatusOK, "我是一个新闻首页")
	})

	r.GET("/TEST", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"massage": "我是一个测试Json界面",
			"data":    http.StatusOK,
		})
	})

	r.GET("/Json1", func(c *gin.Context) {

		article := new(Article)

		article.Title = "Hello world"
		article.Content = "This is a test article wooooow"
		//article := &Article{
		//	Title:   "Hello World",
		//	Content: "This is a test article",
		//}
		c.JSON(http.StatusOK, article)
	})

	r.GET("/josnp", func(c *gin.Context) {
		article := &Article{
			Title:   " JSONP Page",
			Content: "Nothing to see here",
		}
		c.JSONP(http.StatusOK, article)
	})

	// 返回XML数据

	r.GET("/XML", func(c *gin.Context) {
		//article := &Article {
		//	Title : "XML Page ",
		//	Content: "<XML> This is a XML test page </XML>",
		//}
		//
		c.XML(http.StatusOK, gin.H{
			"message": "Hello XML Page",
			"data":    "This is a test XML",
		})

	})

	r.GET("/html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "googleTemplate.html", gin.H{
			"message": "我是一个新闻页面的数据",
			"data":    "This is a news heml",
		})
	})

	// 启动服务
	err := r.Run(":8000")
	if err != nil {
		panic(err)
	}
}
