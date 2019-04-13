package db

import (
	"fmt"

	"../domain"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	db  *gorm.DB
	err error
)

func Init() {
	db, err = gorm.Open("postgres", "host=0.0.0.0 port=5432 user=gorm dbname=gorm password=gorm sslmode=disable")
	fmt.Print(db)
	migrate()
	initMaster()
	if err != nil {
		panic(err)
	}
}

func GetDB() *gorm.DB {
	return db
}

func migrate() {
	db.AutoMigrate(&domain.Order{}, &domain.Currency{})
}

func initMaster() {
	btc := domain.NewCurrency("BTC")
	usd := domain.NewCurrency("USD")
	db.Save(&btc)
	db.Save(&usd)
}
