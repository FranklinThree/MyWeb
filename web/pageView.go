package web

import (
	"gorm.io/gorm"
	"net"
)

// PageView 网页访问记录，仅为PageViewCollector的子目录
type PageView struct {
	gm                   gorm.Model
	UserIP               net.IP
	PageViewCollectorKey string
}
