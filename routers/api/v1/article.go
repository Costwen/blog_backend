package v1

import (
	"blog_backend/models"
	"blog_backend/pkg/e"
	"blog_backend/pkg/util"
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
	err := article.FindByID(id)
	if err != nil {
		Return(http.StatusBadRequest, e.ERROR_NOT_EXIST_ARTICLE, nil, c)
		return
	}
	Return(http.StatusOK, e.SUCCESS, article, c)
}

//获取多个文章
func GetArticles(c *gin.Context) {
	articles := &models.Articles{}
	err := articles.FindByPage(util.Pagination(c))
	if err != nil{
		Return(http.StatusBadRequest, e.INVALID_PARAMS, nil, c)
		return
	}
	Return(http.StatusOK, e.SUCCESS, articles, c)
}

//新增文章
func AddArticle(c *gin.Context) {
	article := &models.Article{}
	err := c.ShouldBind(article)
	if err != nil {
		Return(http.StatusBadRequest, e.INVALID_PARAMS, nil, c)
		return
	}
	result, _ := article.Insert()
	Return(http.StatusOK, e.SUCCESS, result, c)
}

//修改文章
func EditArticle(c *gin.Context) {
	article := &models.Article{}
	article.FindByID(c.Param("id"))
	err := c.ShouldBind(article)
	if err != nil {
		Return(http.StatusBadRequest, e.INVALID_PARAMS, nil, c)
		return 
	}
	result, err := article.Update()
	if err != nil {
		Return(http.StatusBadRequest, e.INVALID_PARAMS, nil, c)
		return 
	}
	Return(http.StatusOK, e.SUCCESS, result, c)
}

//删除文章
func DeleteArticle(c *gin.Context) {
	id, _:= primitive.ObjectIDFromHex(c.Param("id"))
	article := &models.Article{
		Id: id,
	}
	result, _ := article.Delete()
	Return(http.StatusOK, e.SUCCESS, result, c)
}
