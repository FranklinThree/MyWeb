package web

import "net"

type PageViewCollector struct {
	Key       string `gorm:"primary"`
	Count     uint64
	PageViews []*PageView
}

func (pvc *PageViewCollector) AddPageView(ip net.IP) {
	pvc.PageViews[pvc.Count] = &PageView{}
	pvc.PageViews[pvc.Count].UserIP = ip
	pvc.Count++
}
