package service

import "gorm.io/gorm"

var db *gorm.DB

func NewDB(newDB *gorm.DB) {
	db = newDB
}
