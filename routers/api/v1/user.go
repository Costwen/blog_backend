package v1

import (
	"blog_backend/models"
	"blog_backend/pkg/e"
	"blog_backend/pkg/util"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context){
	user := &models.User{}
	err := c.ShouldBind(user)
	if err != nil {
		Return(http.StatusBadRequest, e.INVALID_PARAMS, nil, c)
		return
	}
	islogin := user.Login()
	fmt.Print(islogin)
	if !islogin {
		Return(http.StatusBadRequest, e.ERROR_AUTH, nil, c)
		return
	}
	token, _ := util.GenerateToken(user.Username)
	data := make(map[string]interface{})
	data["token"] = token
	Return(http.StatusOK, e.SUCCESS, data, c)
}