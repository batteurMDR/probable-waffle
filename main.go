package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"louis/pw/cache"
	"louis/pw/db"
	"louis/pw/db/orm"
	"louis/pw/service"
)

type Config struct {
	Env    string
	SQLite struct {
		Name string
	}
}

var config Config

const (
	EnvProduction = "production"
	EnvTest       = "test"
	EnvLocal      = "local"
)

func init() {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")      // optionally look for config in the working directory
	err := viper.ReadInConfig()   // Find and read the config file
	if err != nil {               // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
	log.Println("init viper get ENV value to:", viper.GetString("ENV"))
	config.Env = viper.GetString("ENV")
	config.SQLite.Name = viper.GetString("Sqlite.Name")
}

func main() {
	var dbConnection *db.Store
	switch config.Env {
	case EnvLocal:
		dbConnection = orm.NewSQLite(config.SQLite.Name)
	case EnvProduction:
		dbConnection = orm.NewPostgre("mydb.db")
	}

	router := gin.Default()
	cacheDB := cache.New("", "")

	err := service.New(router, dbConnection, cacheDB)
	if err != nil {
		panic(err)
	}

	router.Run(":" + os.Getenv("PORT"))
}
