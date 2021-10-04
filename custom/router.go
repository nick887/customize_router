package custom

import (
	"github.com/gin-gonic/gin"
	"github.com/nick887/customize_router/adapter"
)

// @author Ritchie
// @version 2.0, 2021/10/4
// @since 2.0.0

// CustmoRouter is the router we define
type CustmoRouter struct {
	gin.IRouter
}
// GETC is a Encapsulation of GET
// Tail processing only the last handler
// processing order exchange
func (cr CustmoRouter) GETC(relativePath string, tailHandler adapter.RequestHandler, frontHandlers ...gin.HandlerFunc) gin.IRoutes {
	if tailHandler == nil {
		return cr.GET(relativePath, frontHandlers...)
	} else {
		frontHandlers = append(frontHandlers, adapter.ErrorAdapter(tailHandler))
		return cr.GET(relativePath, frontHandlers...)
	}
}
// GroupC is a Encapsulation of Group
func (cr CustmoRouter) GroupC(relativePath string, handlers ...gin.HandlerFunc) *CustmoRouter {
	return &CustmoRouter{cr.Group(relativePath, handlers...)}
}

// POSTC is a Encapsulation of POST
func (cr CustmoRouter) POSTC(relativePath string, tailHandler adapter.RequestHandler, frontHandlers ...gin.HandlerFunc) gin.IRoutes  {
	if tailHandler == nil {
		return cr.POST(relativePath, frontHandlers...)
	} else {
		frontHandlers = append(frontHandlers, adapter.ErrorAdapter(tailHandler))
		return cr.POST(relativePath, frontHandlers...)
	}
}

// DELETEC is a Encapsulation of DELETE
func (cr CustmoRouter) DELETEC(relativePath string, tailHandler adapter.RequestHandler, frontHandlers ...gin.HandlerFunc) gin.IRoutes  {
	if tailHandler == nil {
		return cr.DELETE(relativePath, frontHandlers...)
	} else {
		frontHandlers = append(frontHandlers, adapter.ErrorAdapter(tailHandler))
		return cr.DELETE(relativePath, frontHandlers...)
	}
}