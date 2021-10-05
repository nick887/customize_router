package adapter

import (
	"github.com/Chronostasys/centralog/centralog"
	"github.com/gin-gonic/gin"
	"github.com/nick887/customize_router/http_status"
	"net/http"
	"time"
)

// @author Ritchie
// @version 1.0, 2021/10/4
// @since 1.0.0

// RequestHandler is a type which give the router a restriction
type RequestHandler func(ctx *gin.Context) (httpStatus *http_status.HttpStatus, json interface{})

// ResponseError is a unified response of error
type ResponseError struct {
	Time      time.Time   `json:"time"`
	Reason    interface{} `json:"reason"`
	RequestId string      `json:"request_id"`
	Code      int         `json:"code"`
}

// ErrorAdapter is wrapper of handler
func ErrorAdapter(handler RequestHandler) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		httpStatus, json := handler(ctx)
		processResult(ctx, httpStatus, json)
	}
}

// processResult process the error msg
func processResult(ctx *gin.Context, httpStatus *http_status.HttpStatus, json interface{}) {
	if httpStatus.Code != 200 {
		centralog.Error("[request error]").Any("error", httpStatus.Msg).CtxID(ctx).Log()
		if httpStatus.Code >= 1000 {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, generateErrorMsg(ctx, httpStatus))
			return
		}
		ctx.AbortWithStatusJSON(httpStatus.Code, generateErrorMsg(ctx, httpStatus))
		return
	}
	ctx.JSON(httpStatus.Code, json)
}

// generateErrorMsg generate the error msg
func generateErrorMsg(ctx *gin.Context, httpStatus *http_status.HttpStatus) ResponseError {
	requestId, _ := ctx.Get(centralog.IDKey)
	return ResponseError{
		Time:      time.Now(),
		Reason:    httpStatus.Msg,
		RequestId: requestId.(string),
		Code:      httpStatus.Code,
	}
}
