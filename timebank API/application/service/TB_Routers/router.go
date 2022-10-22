package TB_Routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	v1 "timebank/application/service/TB_Api/v1"
	"timebank/application/service/middlewares"
)

func InitRouter(service *v1.ServiceSetup) *gin.Engine {
	r := gin.Default()
	r.Use(middlewares.Cors())
	apiV1 := r.Group("/api/v1")
	{
		apiV1.GET("/hello", v1.Hello)
		apiV1.POST("/AcceptServicing", service.AcceptServicing)
		apiV1.POST("/CloseServicing", service.CloseServicing)
		apiV1.POST("/CreateJobs", service.CreateJobs)
		apiV1.POST("/CreateOrgs", service.CreateOrgs)
		apiV1.POST("/CreateServicing", service.CreateServicing)
		apiV1.POST("/CreateUsers", service.CreateUsers)
		apiV1.POST("/DoneServicing", service.DoneServicing)
		apiV1.POST("/GetUpdateHistoryList", service.GetUpdateHistoryList)
		apiV1.POST("/InheritAsset", service.InheritAsset)
		apiV1.POST("/JobList", service.JobList)
		apiV1.POST("/ManagerList", service.ManagerList)
		apiV1.POST("/OrgList", service.OrgList)
		apiV1.POST("/RechargeAsset", service.RechargeAsset)
		apiV1.POST("/RechargeList", service.RechargeList)
		apiV1.POST("/ServicingStatusList", service.ServicingStatusList)
		apiV1.POST("/SpecialTxList", service.SpecialTxList)
		apiV1.POST("/TransferAsset", service.TransferAsset)
		apiV1.POST("/TxList", service.TxList)
		apiV1.POST("/UpdateServiceInfo", service.UpdateServiceInfo)
		apiV1.POST("/UpdateUserInfo", service.UpdateUserInfo)
		apiV1.POST("/UserList", service.UserList)
	}
	r.StaticFS("/web", http.Dir("./dist/"))
	return r
}
