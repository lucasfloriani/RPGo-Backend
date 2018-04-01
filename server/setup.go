package server

import (
	"rpgo-backend/middleware"
	"rpgo-backend/router"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func Setup(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(middleware.DBtoContext(db))
	router.InitializeRoutes(r)
	return r
}
