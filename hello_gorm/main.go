package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

type User struct {
	ID 			 uint	 `gorm:"primary_key;unique_index;column:id"`
	Name         string  `gorm:"type:varchar(100);column:name"`
	Email        string  `gorm:"type:varchar(100);column:email"`
}

func (User) TableName() string {
	return "profiles"
}

func main() {
	db, err := gorm.Open("mysql", "root:root@/hello_gorm_db?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()

	if err != nil {
		fmt.Printf("error to open connect: %v", err)
	}

	db.CreateTable(&User{})

	db.CreateTable(&RegistryEntity{})

	user := User{Name: "Daniel Acaz", Email: "daniel_acaz@hotmail.com"}

	registry := RegistryEntity{Date: time.Now(), Category: "categoryTest", Title: "titleTest", Amount: 100.0,
		MyCategory: "myCategoryTest", FamilyCategory: "familyCategoryTest"}

	db.Create(&user)
	db.Create(&registry)

}

type RegistryEntity struct {
	ID             int64     `gorm:"primary_key;AUTO_INCREMENT;column:id"`
	Date           time.Time `gorm:"column:date"`
	Category       string    `gorm:"type:varchar(255);column:category"`
	Title          string    `gorm:"type:varchar(255);not null;column:title"`
	Amount         float64   `gorm:"type:varchar(255);not null;column:amount"`
	MyCategory     string    `gorm:"type:varchar(255);column:my_category"`
	FamilyCategory string    `gorm:"type:varchar(255);column:family_category"`
}

func (RegistryEntity) TableName() string {
	return "registry"
}
