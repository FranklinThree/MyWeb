package web

import "net"

// PageViewCollector
type PageViewCollector struct {
	Key       string `gorm:"primarykey"`
	Count     uint64
	PageViews []*PageView
}

func (pvc *PageViewCollector) AddPageView(ip net.IP) {
	pvc.PageViews[pvc.Count] = &PageView{}
	pvc.PageViews[pvc.Count].UserIP = ip
	pvc.Count++
}
