package main

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Todo struct {
	gorm.Model
	Name   string `json:"name" binding:"required"`
	Title  string `json:"title" binding:"required"`
	UserId uint   `json:"-" binding:"required"`
}
