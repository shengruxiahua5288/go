package db

import (
	"github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
	"log"
)

type BaseGorm struct {
	DB *gorm.DB
}

var baseDB *gorm.DB

func (this *BaseGorm) InitDB() {
	var err error
	this.DB, err = gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/mydb?charset=utf8&parseTime=true&loc=Local")
	if err != nil {
		log.Fatal(err)
		return
	}
	this.DB.SingularTable(true)
	this.DB.DB().SetMaxIdleConns(10)
	this.DB.DB().SetMaxOpenConns(100)
	this.DB.DB().SetConnMaxLifetime(300*time.Second)
	this.DB.LogMode(true)
	baseDB = this.DB
}
func (this *BaseGorm) GetDB() (DB *gorm.DB) {
	if baseDB != nil {
		DB = baseDB
	} else {
		log.Fatal("DB not initial.")
		return
	}
	return
}
