package Models

import (
	"fmt"
	"practice1servergo/Database"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Cui        string `json:"cui"`
	Names      string `json:"names"`
	Lastnames  string `json:"lastnames"`
	Birth_date string `json:"birth_date"`
	Password   []byte `json:"password"`
	Rol_id     int    `json:"rol_id"`
	Rol        Rol    `gorm:"foreignkey:Rol_id"`
}

func (b *User) TableName() string {
	return "user"
}

func GetAllUser(b *[]User) (err error) {
	if err = Database.DB.Find(b).Error; err != nil {
		return err
	}
	return nil
}

func AddNewUser(b *User) (err error) {
	if err = Database.DB.Create(b).Error; err != nil {
		return err
	}
	return nil
}

func GetOneUser(b *User, id string) (err error) {
	if err := Database.DB.Where("id = ?", id).First(b).Error; err != nil {
		return err
	}
	return nil
}

func PutOneUser(b *User, id string) (err error) {
	fmt.Println(b)
	Database.DB.Save(b)
	return nil
}

func DeleteUser(b *User, id string) (err error) {
	Database.DB.Where("id = ?", id).Delete(b)
	return nil
}
