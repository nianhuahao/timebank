package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"timebank/application/service/TB_Utils"
)

func Hello(c *gin.Context) {
	respG := TB_Utils.Gin{C: c}
	respG.Response(http.StatusOK, "成功", map[string]interface{}{
		"msg": "Hello",
	})
}
