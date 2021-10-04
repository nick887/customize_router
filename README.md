# Customize router

## it is a encapsulation of gin router

## example

```go
func main() {
    r := gin.Default()
    cr := &custom.CustmoRouter{r}
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

func Hello(ctx *gin.Context,request interface{}) (status int, json interface{}, err error) {
    return http.StatusOK, gin.H{
    "msg": "hello",
    }, nil
}
```

## details
1. init centralog first
