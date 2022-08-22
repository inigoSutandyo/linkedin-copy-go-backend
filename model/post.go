package model

import (
	"fmt"

	"github.com/inigoSutandyo/linkedin-copy-go/utils"
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title      string `json:"-"`
	Content    string `json:"content" gorm:"text"`
	Attachment string `json:"attachment"`
	Likes      int    `json:"like"`
	UserID     uint
	User       User
	// Template   Template `gorm:"embedded"`s
	// PostLikes  []PostLike
}

func CreatePost(user *User, post *Post) error {
	fmt.Println(user.Email)
	fmt.Println(post.Content)
	err := utils.DB.Model(user).Association("Posts").Append(post)
	return err
}

func GetAllPost(posts *[]Post, users *[]User) error {
	// err := utils.DB.Find(posts).Error
	err := utils.DB.Preload("User").Find(posts).Error
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(posts)
	}
	return err
}
