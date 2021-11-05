package main

import (
	"gorm.io/gorm"
)

type PageView struct {
	gm                   gorm.Model
	PageViewCollectorKey string
}
