package migrations

import (
	models "github.com/inigoSutandyo/linkedin-copy-go/model"
	utils "github.com/inigoSutandyo/linkedin-copy-go/utils"
)

func Migrate() {
	// utils.DB.Migrator().CreateConstraint(&models.User{}, "Posts")
	// utils.DB.Migrator().DropTable(&models.Post{})
	// utils.DB.Migrator().DropTable(&models.Comment{})
	utils.DB.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{}, &models.Reply{})
}
