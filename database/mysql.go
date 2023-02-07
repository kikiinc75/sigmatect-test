package database

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	dbUser     = "root"
	dbPassword = "root"
	dbHost     = "localhost"
	dbPort     = 3306
	dbName     = "sigmatect_db"
)

func ConnectDB() *gorm.DB {
	dsn := os.Getenv("MYSQL_DSN")
	if dsn == "" {
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", dbUser, dbPassword, dbHost, dbPort, dbName)
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}
