package util

import (
	"blog_backend/pkg/setting"

	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Pagination(c *gin.Context) (*options.FindOptions) {
	page := com.StrTo(c.DefaultQuery("page", "0")).MustInt()
	if page < 0 {
		page = 0
	}
	options := options.Find()
	start := page * setting.PageSize
	options.SetSkip(int64(start))
	options.SetLimit(int64(setting.PageSize))

	return options
}
