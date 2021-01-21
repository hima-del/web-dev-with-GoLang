package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type product struct {
	gorm.Model
	code  string
	price int
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&product)
	db.Create(&product{code: "d42", price: 100})
	var product Product
	//find product with integer primary key
	db.First(&product, 1)
	//find product with code D42
	db.First(&product, "code=?", "d42")
	// Update - update product's price to 200
	db.Model(&product).Update("price",200)
   //update multiple fields
   db.Model(&product).Updates(Product{price:200,code:"f42"})
   db.Model(&product).Updates(map[string]interface{}{"price:200","code":"f42"})
   //delete product
   db.Delete(&product,1)
}
