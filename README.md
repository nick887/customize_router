# Customize router

## it is a encapsulation of gin router

## demo

```go
func main() {
    r := gin.Default()
    cr := &custom.CustmoRouter{IRouter: r}
    customRclt.InitCustmoR(cr)
	UseContentRouter(cr)
	_ = r.Run(":9999")
}


func UseContentRouter(r *custom.CustmoRouter) {
    a := r.GroupC("/api")
    {
        hole := a.GroupC("/hole")
        {
        // the second handler must be type of RequestHandler
        // handlers after the second must be type of gin.HandlerFunc
        // gin.HandlerFunc will excute first
        hole.GETC("/hello",&HelloRequest{} ,Hello)
        }
    }
}

func Hello(ctx *gin.Context) (statusCode *http_status.HttpStatus, json interface{}) {
    return &http_status.HttpStatus{
    Msg:  "error",
	// custmo http error code
	// it always return 500
	// code msg will show in msg
    Code: common.HOLE_NOT_FOUND,
    }, gin.H{
    "msg": "hello",
    }
}

func InitCustmoR(cr *custom.CustmoRouter) {
    logSettings:=config.Conf.LogSettings
    customize_router.InitCustomRouter(&centralog.LogOptions{
    Server:     logSettings.LogCenterHost,
    Db:         logSettings.Db,
    Collection: logSettings.Collection,
    },cr)
}
```

## details
1. init centralog first