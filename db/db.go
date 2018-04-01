package db

import (
	"flag"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/configor"
	"github.com/jinzhu/gorm"
	"github.com/qor/validations"

	_ "github.com/jinzhu/gorm/dialects/mssql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

/**
 * Config is loaded from configuration file
 */
var Config = struct {
	Test struct {
		Adapter  string `required:"true"`
		Database string `required:"true"`
	}

	Production struct {
		Adapter  string `required:"true"`
		Database string `required:"true"`
	}
}{}

var (
	Migrate string
	ShowSql string
)

/**
 * Init does
 */
func init() {
	Migrate = os.Getenv("MIGRATE")
	flag.StringVar(&Migrate, "m", Migrate, "AutoMigrate")
	flag.StringVar(&ShowSql, "s", ShowSql, "ShowSql")
	flag.Parse()
}

/**
 * Connect does
 */
func Connect() *gorm.DB {
	configor.Load(&Config, "config.yml")

	db, err := gorm.Open(Config.Production.Adapter, Config.Production.Database)
	if err != nil {
		log.Fatalf("Got error when connect database, the error is '%v'", err)
	}
	validations.RegisterCallbacks(db)

	flagOptions(db)

	return db
}

/**
 * flagOptions does
 */
func flagOptions(db *gorm.DB) {
	db.LogMode(false)
	if ShowSql != "" {
		db.LogMode(true)
	}

	if Migrate == "1" {
		migrate(db)
	}
}

/**
 * migrate does
 */
func migrate(db *gorm.DB) {
	db.DropTableIfExists()

	db.AutoMigrate()
}

/**
 * DBInstance does
 */
func DBInstance(c *gin.Context) *gorm.DB {
	return c.MustGet("DB").(*gorm.DB)
}
