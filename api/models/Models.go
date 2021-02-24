/*
 * @Time    : 2020年11月23日 09:52:11
 * @Author  : root
 * @Project : micro
 * @File    : key.go
 * @Software: GoLand
 * @Describe:
 */
package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name   string `json:"name"`
	email   string `json:"email"`
	Image   string `json:"image"`
	PassWord    string `json:"pass_word"`
}

type Post struct {
	gorm.Model
	Title   string `json:"title"`
	Image   string `json:"image"`
	Type    string `json:"type"`
	Tags    string `json:"tags"`
	View    int    `json:"view"`
	Content string `json:"content" gorm:"type:text;not null"`
}

func (k *Post) BeforeSave() {
	k.View = 0
}

type Tag struct {
	gorm.Model
	Title string `json:"title"`
}

type Comment struct {
	gorm.Model
	PostId    string `json:"post_id"`
	UserName  string `json:"user_name"`
	UserImage string `json:"user_image"`
	RootId    string `json:"root_id"`
	Content   string `json:"content" gorm:"type:text;not null"`
}

type Link struct {
	gorm.Model
	Title   string `json:"title"`
	Url     string `json:"url"`
	Content string `json:"content" gorm:"type:text;not null"`
}
