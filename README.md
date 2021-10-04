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
        hole.GETC("/hello", Hello)
        }
    }
}

func Hello(ctx *gin.Context) (status int, json interface{}, err error) {
    return http.StatusOK, gin.H{
    "msg": "hello",
    }, nil
}
```

## details
1. init centralog first
