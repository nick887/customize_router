package adapter

import (
	"github.com/Chronostasys/centralog/centralog"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// @author Ritchie
// @version 2.0, 2021/10/4
// @since 2.0.0

// RequestHandler is a type which give the router a restriction
type RequestHandler func(ctx *gin.Context,request interface{}) (status int, json interface{}, err error)

// ResponseError is a unified response of error
type ResponseError struct {
	Time      time.Time `json:"time"`
	Reason    string    `json:"reason"`
	RequestId string    `json:"request_id"`
}

// ErrorAdapter is wrapper of handler
func ErrorAdapter(t interface{}, handler RequestHandler) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// need validation
		if t != nil{
			// validation
			// fail
			if err := ctx.ShouldBind(t); err != nil {
				processResult(ctx,http.StatusBadRequest,nil,err)
				return
			}
			// success
		}
		status, json, err := handler(ctx,t)
		processResult(ctx, status, json, err)
	}
}

// processResult process the error msg
func processResult(ctx *gin.Context, status int, json interface{}, err error) {
	if err != nil {
		centralog.Error("[request error]").Any("error", err).CtxID(ctx).Log()
		ctx.AbortWithStatusJSON(status, generateErrorMsg(ctx, err))
		return
	}
	ctx.JSON(status, json)
}

// generateErrorMsg generate the error msg
func generateErrorMsg(ctx *gin.Context, err error) ResponseError {
	requestId, _ := ctx.Get(centralog.IDKey)
	return ResponseError{
		Time:      time.Now(),
		Reason:    err.Error(),
		RequestId: requestId.(string),
	}
}
