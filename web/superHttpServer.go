package web

import (
	"github.com/gin-gonic/gin"
	"net"
)

type SuperHttpServer struct {
	*gin.Engine
	SiteNames map[string]*PageViewCollector
}

func (shs *SuperHttpServer) GET(relativePath string, handlers func(c *gin.Context)) {
	siteName, NotNil := shs.SiteNames[relativePath]
	if !NotNil {
		shs.SiteNames[relativePath] = &PageViewCollector{}
		shs.SiteNames[relativePath].Key = relativePath
	}
	shs.GET(relativePath, func(c *gin.Context) {
		ip, isOk := c.RemoteIP()
		if isOk {
			siteName.AddPageView(ip)
		} else {
			siteName.AddPageView(net.IP{1, 1, 1, 1})
		}
		handlers(c)
	})
}
func (shs *SuperHttpServer) POST(relativePath string, handlers func(c *gin.Context)) {
	siteName, NotNil := shs.SiteNames[relativePath]
	if !NotNil {
		shs.SiteNames[relativePath] = &PageViewCollector{}
		shs.SiteNames[relativePath].Key = relativePath
	}
	shs.POST(relativePath, func(c *gin.Context) {
		ip, isOk := c.RemoteIP()
		if isOk {
			siteName.AddPageView(ip)
		} else {
			siteName.AddPageView(net.IP{1, 1, 1, 1})
		}
		handlers(c)
	})
}
