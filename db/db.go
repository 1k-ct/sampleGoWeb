package db

import (
	"github.com/jinzhu/gorm"

	"github.com/1k-ct/sampleGoWeb/entity" //自作パッケージ　githubからimport ("go get github.com/username/appname/...")
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db  *gorm.DB
	err error
)

// Init is initialize db from main function
func Init() {
	db, err = gorm.Open("mysql", "user1:Password_01@/go_sample?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	autoMigration()
}

// GetDB is called in models
func GetDB() *gorm.DB {
	return db
}

// Close is closing db
func Close() {
	if err := db.Close(); err != nil {
		panic(err)
	}
}
func autoMigration() {
	db.AutoMigrate(&entity.User{})
}
