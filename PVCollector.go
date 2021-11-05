package main

type PageViewCollector struct {
	key       string `gorm:"primarykey"`
	count     uint64
	PageViews []PageView
}
