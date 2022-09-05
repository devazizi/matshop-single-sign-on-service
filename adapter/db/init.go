package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DB struct {
	store *gorm.DB
}

func NewDB(dsn string) DB {
	connection, dbConnectionErr := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if dbConnectionErr != nil {
		panic("can not connect to db")
	}

	var entities []any = []any{
		&User{},
		&OauthAccessToken{},
	}

	if migrationErr := connection.AutoMigrate(entities...); migrationErr != nil {
		panic("auto migration fail")
	}

	return DB{
		store: connection,
	}
}
