package services

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type CmsApp struct {
	db  *gorm.DB
	rdb *redis.Client
}

func NewCmsApp() *CmsApp {
	app := &CmsApp{}

	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	connDB(app)
	connRdb(app)

	return app
}

func connDB(app *CmsApp) {
	dbUrl := os.Getenv("DB_URL")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbUrl)
	mysqlDB, err := gorm.Open(mysql.Open(dsn))

	if err != nil {
		panic(err)
	}

	db, err := mysqlDB.DB()

	if err != nil {
		panic(err)
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)

	app.db = mysqlDB
}

func connRdb(app *CmsApp) {
	redisDBStr := os.Getenv("REDIS_DB")
	redisDB, err := strconv.Atoi(redisDBStr)

	if err != nil {
		log.Fatalf("Error converting REDIS_DB to int: %v", err)
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       redisDB,
	})

	_, err = rdb.Ping(context.Background()).Result()

	if err != nil {
		panic(err)
	}

	app.rdb = rdb
}
