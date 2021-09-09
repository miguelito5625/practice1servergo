package Models

import (
	"fmt"
	"practice1servergo/Database"

	"gorm.io/gorm"
)

type Usuario struct {
	gorm.Model
	Cui        string `json:"cui"`
	Nombres    string `json:"nombres"`
	Apeliidos  string `json:"apellidos"`
	Nacimiento string `json:"nacimiento"`
	Clave      string `json:"clave"`
	Rol_id     int    `json:"rol_id"`
	Rol        Rol    `gorm:"foreignkey:Rol_id"`
}

func (b *Usuario) TableName() string {
	return "usuario"
}

func GetAllUsuario(b *[]Usuario) (err error) {
	if err = Database.DB.Joins("Rol").Find(b).Error; err != nil {
		return err
	}
	return nil
}

func AddNewUsuario(b *Usuario) (err error) {
	if err = Database.DB.Create(b).Error; err != nil {
		return err
	}
	return nil
}

func GetOneUsuario(b *Usuario, id string) (err error) {
	if err := Database.DB.Where("id = ?", id).First(b).Error; err != nil {
		return err
	}
	return nil
}

func LoginUsuario(b *Usuario, cui string) (err error) {
	if err := Database.DB.Where("cui = ?", cui).First(b).Error; err != nil {
		return err
	}
	return nil
}

func PutOneUsuario(b *Usuario, id string) (err error) {
	fmt.Println(b)
	Database.DB.Save(b)
	return nil
}

func DeleteUsuario(b *Usuario, id string) (err error) {
	Database.DB.Where("id = ?", id).Delete(b)
	return nil
}
