package routers

import (
	"blog_backend/middleware/jwt"
	"blog_backend/pkg/setting"
	v1 "blog_backend/routers/api/v1"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {

	gin.SetMode(setting.RunMode)

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	apiv1 := r.Group("/api/v1")
	{
		apiv1.GET("/articles", v1.GetArticles)
		//获取指定文章
		apiv1.GET("/articles/:id", v1.GetArticle)
	}
	apiv1.Use(jwt.JWT())
	{
		apiv1.POST("/articles", v1.AddArticle)
		//更新指定文章
		apiv1.PATCH("/articles/:id", v1.EditArticle)
		//删除指定文章
		apiv1.DELETE("/articles/:id", v1.DeleteArticle)
	}
	return r
}
