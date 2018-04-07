package db

import (
	"flag"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/configor"
	"github.com/jinzhu/gorm"

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

	Dev struct {
		Adapter  string `required:"true"`
		Database string `required:"true"`
	}

	Production struct {
		Adapter  string `required:"true"`
		Database string `required:"true"`
	}
}{}

var (
	Migrate  string
	ShowSql  string
	DataBase string
)

/**
 * Init does
 */
func init() {
	Migrate = os.Getenv("MIGRATE")
	flag.StringVar(&Migrate, "m", Migrate, "AutoMigrate")
	flag.StringVar(&ShowSql, "s", ShowSql, "ShowSql")
	flag.StringVar(&DataBase, "db", DataBase, "DataBase")
	flag.Parse()
}

/**
 * Connect does
 */
func Connect() *gorm.DB {
	configor.Load(&Config, "config.yml")

	db, err := getEnvDatabase()
	if err != nil {
		log.Fatalf("Got error when connect database, the error is '%v'", err)
	}

	flagOptions(db)

	return db
}

/**
 * DBInstance does
 */
func DBInstance(c *gin.Context) *gorm.DB {
	return c.MustGet("DB").(*gorm.DB)
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

func getEnvDatabase() (*gorm.DB, error) {
	switch DataBase {
	case "test":
		return gorm.Open(Config.Test.Adapter, Config.Test.Database)
	case "dev":
		return gorm.Open(Config.Dev.Adapter, Config.Dev.Database)
	default:
		return gorm.Open(Config.Production.Adapter, Config.Production.Database)
	}
}
