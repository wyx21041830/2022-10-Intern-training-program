package main
import (
    "fmt"
    "github.com/labstack/echo/v4"
    "net/http"
    "github.com/sirupsen/logrus"
)
requestLogger:= logrus.WithFields(logrus.Fields{user:"sss"})//log日志的设置
func handlePost(c echo.Context) error {
    return c.String(http.StatusOK, "pong!")
}

func handlePost(c echo.Context)error{
    query,err:=c.QueryParam("query")
    if err!=nil{
        fmt.Println(err)
        requestLogger.Error("error msg")
    }
    return c.String(http.StatusOK, query)
}
func handleRequet(c echo.Context)error{
    var bodyBytes []byte
    defer c.Request().Body.Close()
    bodyBytes,err:=ioutil.ReadAll(c.Request().Body)
    if err!=nil{
        fmt.Println("requset error",err)
        requestLogger.Error("error msg")
    }
    var info string
    if err =json.Unmarshal(bodyBytes, &info);err!=nil{
        fmt.Println("error",err)
        requestLogger.Error("error msg")
        return c.String(http.StatusOK,err)
    }
    return c.String(http.StatusOK,info)//将读取的数据按照字符串形式输出
}  
type response struct {
    code uint
    msg string
    data interface{}
}
func Response(c echo.Context, code int, msg string, data ...interface{}){
    return c.JSON(http.StatusOK, Response{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}
func main() {
    e := echo.New()
    e.GET("/pong", handlePost)//路由
    e.POST("/print/query",handlePost)//请求
    e.POST("/print/body",handleBodyd)//请求
    if err:=e.Start("127.0.0.1:80");err!=nil{//启动
        fmt.Println("echo start error",err)
        return
    }
    
}

