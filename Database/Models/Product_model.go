package Models

import (
	"fmt"
	"practice1servergo/Database"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name  string  `json:"name"`
	Code  string  `json:"code"`
	Price float32 `json:"price"`
}

func (b *Product) TableName() string {
	return "product"
}

func GetAllProduct(b *[]Product) (err error) {
	if err = Database.DB.Find(b).Error; err != nil {
		return err
	}
	return nil
}

func AddNewProduct(b *Product) (err error) {
	if err = Database.DB.Create(b).Error; err != nil {
		return err
	}
	return nil
}

func GetOneBook(b *Product, id string) (err error) {
	if err := Database.DB.Where("id = ?", id).First(b).Error; err != nil {
		return err
	}
	return nil
}

func PutOneBook(b *Product, id string) (err error) {
	fmt.Println(b)
	Database.DB.Save(b)
	return nil
}

func DeleteBook(b *Product, id string) (err error) {
	Database.DB.Where("id = ?", id).Delete(b)
	return nil
}
