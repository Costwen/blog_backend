package v1

import (
	"blog_backend/models"
	"blog_backend/pkg/e"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Return(status int, code int, data interface{}, c *gin.Context) {
	c.JSON(status, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

//获取单个文章
func GetArticle(c *gin.Context) {
	article := &models.Article{}
	id := c.Param("id")
	code := e.SUCCESS
	err := article.FindByID(id)
	if err != nil {
		code = e.ERROR_NOT_EXIST_ARTICLE
		Return(http.StatusBadRequest, code, nil, c)
		return
	}
	Return(http.StatusOK, code, article, c)
}

//获取多个文章
func GetArticles(c *gin.Context) {

}

//新增文章
func AddArticle(c *gin.Context) {
	code := e.SUCCESS
	article := &models.Article{}
	err := c.ShouldBind(article)
	if err != nil {
		code = e.INVALID_PARAMS
		Return(http.StatusBadRequest, code, nil, c)
		return
	}
	result, _ := article.Insert()
	Return(http.StatusOK, code, result, c)
}

//修改文章
func EditArticle(c *gin.Context) {
	article := &models.Article{}
	c.Bind(&article)
	fmt.Print(article)
}

//删除文章
func DeleteArticle(c *gin.Context) {
	id, _:= primitive.ObjectIDFromHex(c.Param("id"))
	article := &models.Article{
		Id: id,
	}
	result, _ := article.Delete()
	code := e.SUCCESS
	Return(http.StatusOK, code, result, c)
}
