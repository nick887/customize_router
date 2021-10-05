package log

import (
	"fmt"
	"github.com/Chronostasys/centralog/centralog"
	"github.com/gin-gonic/gin"
	"github.com/nick887/customize_router/common"
	"time"
	"github.com/google/uuid"
)

// @author Ritchie
// @version 1.0, 2021/10/5
// @since 1.0.0


func Log() gin.HandlerFunc {
	return func(c *gin.Context) {
		// request id
		requestId := c.GetHeader(common.REQUEST_ID)
		var id string
		if requestId == "" {
			id = uuid.New().String()
			c.Header(common.REQUEST_ID, id)
		} else {
			id = c.GetHeader(common.REQUEST_ID)
		}
		method := c.Request.Method
		url := c.Request.URL.RequestURI()
		c.Set(centralog.IDKey, id)
		defer centralog.Sync()
		start := time.Now()
		c.Next()
		centralog.Info(method).
			Any("elapsed",
				fmt.Sprintf("%f", float64(time.Since(start).Microseconds())/1000)+"ms").
			Any("url", url).
			Any("ip", c.ClientIP()).
			CtxID(c).
			Log()
	}
}
