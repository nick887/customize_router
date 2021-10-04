package custom

import (
	"customize_router/adapter"
	"github.com/gin-gonic/gin"
)

// @author Ritchie
// @version 2.0, 2021/10/4
// @since 2.0.0

type CustmoRouter struct {
	gin.IRouter
}

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
func (cr CustmoRouter) GroupC(relativePath string, handlers ...gin.HandlerFunc) *CustmoRouter {
	return &CustmoRouter{cr.Group(relativePath, handlers...)}
}

func (cr CustmoRouter) POSTC(relativePath string, tailHandler adapter.RequestHandler, frontHandlers ...gin.HandlerFunc) gin.IRoutes  {
	if tailHandler == nil {
		return cr.POST(relativePath, frontHandlers...)
	} else {
		frontHandlers = append(frontHandlers, adapter.ErrorAdapter(tailHandler))
		return cr.POST(relativePath, frontHandlers...)
	}
}

func (cr CustmoRouter) DELETEC(relativePath string, tailHandler adapter.RequestHandler, frontHandlers ...gin.HandlerFunc) gin.IRoutes  {
	if tailHandler == nil {
		return cr.DELETE(relativePath, frontHandlers...)
	} else {
		frontHandlers = append(frontHandlers, adapter.ErrorAdapter(tailHandler))
		return cr.DELETE(relativePath, frontHandlers...)
	}
}