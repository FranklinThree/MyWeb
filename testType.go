package main

import "gorm.io/gorm"

type Test struct {
	gorm.Model
	Sentence string `gorm:"default null"`
}
