package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Database *gorm.DB

var uriDSN = "root:!@!@soumya098@tcp(localhost:3306)/mydatabase"
var err error

func DataMigration() {
	Database, err = gorm.Open(mysql.Open(uriDSN), &gorm.Config{})
	if err != nil {
		fmt.Print(err.Error())
		panic("connection failed :(")
	}
	Database.AutoMigrate(&Student{})
}
