package main

import (
	"fmt"
	"os"
	v1 "timebank/application/service/TB_Api/v1"
	"timebank/application/service/TB_Routers"
	"timebank/application/service/TB_Sdk"
)

func main() {

	// sdkInit setup
	sdk, err := TB_Sdk.Setup("config.yaml", &TB_Sdk.Info)
	if err != nil {
		fmt.Println(">> SDK setup error:", err)
		os.Exit(-1)
	}

	// create channel and join
	if err := TB_Sdk.CreateAndJoinChannel(&TB_Sdk.Info); err != nil {
		fmt.Println(">> Create channel and join error:", err)
		os.Exit(-1)
	}

	// create chaincode lifecycle
	if err := TB_Sdk.CreateCCLifecycle(&TB_Sdk.Info, 1, false, sdk); err != nil {
		fmt.Println(">> create chaincode lifecycle error:", err)
		os.Exit(-1)
	}

	// invoke chaincode set status
	fmt.Println(">> 通过链码外部服务设置链码状态......")
	fmt.Println("----------------------------------------------")
	// 初始化服务
	//var service []*TB_Sdk.ServiceSetup
	//for i := 0; i < len(TB_Sdk.Info.Orgs); i++ {
	service, err := v1.InitService(TB_Sdk.Info.ChaincodeID, TB_Sdk.Info.ChannelID, TB_Sdk.Info.Orgs[0], sdk)
	//	if err != nil {
	//		fmt.Printf(">> 第%d个服务端初始化错误 !!!!!!\n", i)
	//	}
	//}
	//var initedUser TB_Models.User

	//r := gin.Default()
	//urlChoose := TB_Routers.InitRouter(r, service[0])
	//urlChoose.Run(":8410")
	r := TB_Routers.InitRouter(service)
	r.Run(":8410")
}
