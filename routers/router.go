package routers

import (
	// jwt "blog_backend/middleware"
	"blog_backend/pkg/setting"
	v1 "blog_backend/routers/api/v1"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {

	gin.SetMode(setting.RunMode)

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	apiv1 := r.Group("/v1")
	{
		apiv1.POST("/login", v1.Login)
		apiv1.GET("/article", v1.GetArticles) // 分页获取
		apiv1.GET("/article/:id", v1.GetArticle)
		sapiv1 := apiv1.Group("/auth")
		// sapiv1.Use(jwt.JWT())
		{
			sapiv1.POST("/article", v1.AddArticle)
			sapiv1.PATCH("/article/:id", v1.EditArticle)
			sapiv1.DELETE("/article/:id", v1.DeleteArticle)
		}
	}

	return r
}
