package Migrations

import (
	"practice1servergo/Database"
	"practice1servergo/Database/Models"
)

func MigrateAll() {
	Database.DB.AutoMigrate(&Models.Rol{})
	Database.DB.AutoMigrate(&Models.User{})
}
