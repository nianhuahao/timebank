package TB_Utils

import "github.com/gin-gonic/gin"

type Gin struct {
	C *gin.Context
}

type Response struct {
	Code int
	Msg  string
	Data interface{}
}

func (g *Gin) Response(httpCode int, errMsg string, data interface{}) {
	g.C.JSON(httpCode, Response{
		Code: httpCode, // 报错状态
		Msg:  errMsg,   // 报错
		Data: data,     // 报错详细
	})
	return
}
