package internal

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Star struct {
	gorm.Model
	Taskname string
	Stars    int
	Status   bool
}

type Dailystar struct {
	gorm.Model
	Taskname string
	Stars    int
}

func Dbde() gorm.DB {
	dsn := "host=127.0.0.1 user=postgres password=admin dbname=star port=5432 sslmode=disable"
	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db.AutoMigrate(&Star{})
	return *db
}

func Dbdaily() gorm.DB {
	dsn := "host=127.0.0.1 user=postgres password=admin dbname=star port=5432 sslmode=disable"
	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db.AutoMigrate(&Dailystar{})
	return *db
}
