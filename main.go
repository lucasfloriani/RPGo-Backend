package main

import (
	"flag"
	"os"
	db "rpgo-backend/db"
	"rpgo-backend/server"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var (
	Port  string
	Debug string
)

func init() {
	Port = os.Getenv("PORT")
	Debug = os.Getenv("DEBUG")
	flag.StringVar(&Port, "p", Port, "Port")
	flag.StringVar(&Debug, "d", Debug, "DebugMode")
}

/*main does*/
func main() {
	Db := db.Connect()
	defer Db.Close()

	router, port := flagOptions(Db)

	router.Run(":" + port)
}

/*flagOptions*/
func flagOptions(Db *gorm.DB) (router *gin.Engine, port string) {
	if Debug == "" {
		gin.SetMode(gin.ReleaseMode)
	}
	router = server.Setup(Db)

	port = "8080"
	if Port != "" {
		if _, err := strconv.Atoi(Port); err == nil {
			port = Port
		}
	}

	return router, port
}
