module github.com/Costwen/blog_backend

go 1.16

require (
	github.com/gin-gonic/gin v1.7.7
	github.com/go-ini/ini v1.66.2
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/jinzhu/gorm v1.9.16 // indirect
	github.com/mattn/go-sqlite3 v2.0.1+incompatible // indirect
	github.com/stretchr/testify v1.7.0 // indirect
	github.com/unknwon/com v1.0.1
)

replace (
	github.com/Costwen/blog_backend/middleware => ./blog_backend/middleware
	github.com/Costwen/blog_backend/models => ./blog_backend/models
	github.com/Costwen/blog_backend/pkg/e => ./blog_backend/pkg/e
	github.com/Costwen/blog_backend/pkg/setting => ./blog_backend/pkg/setting
	github.com/Costwen/blog_backend/pkg/util => ./blog_backend/pkg/util
	github.com/Costwen/blog_backend/routers => ./blog_backend/routers
)
