package TB_Sdk

//
//import (
//	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
//)
//
////type ServiceSetup struct {
////	ChaincodeID string
////	Client      *channel.Client
////}
////
////var t *ServiceSetup
//
//func ChannelExecute(fcn string, args [][]byte) (channel.Response, error) {
//	// 创建客户端，表明在通道的身份
//	//ctx := sdk.ChannelContext(TB_conf.Info.ChannelID, fabsdk.WithUser("Admin"))
//	//cli, err := channel.New(ctx)
//	//if err != nil {
//	//	return channel.Response{}, err
//	//}
//	// 对区块链账本的写操作（调用了链码的invoke）
//	resp, err := t.Client.Execute(channel.Request{
//		ChaincodeID: CC_Name,
//		Fcn:         fcn,
//		Args:        args,
//	})
//	if err != nil {
//		return channel.Response{}, err
//	}
//	//返回链码执行后的结果
//	return resp, nil
//}
//
//// ChannelQuery 区块链查询
//func ChannelQuery(fcn string, args [][]byte) (channel.Response, error) {
//	// 创建客户端，表明在通道的身份
//	//ctx := sdk.ChannelContext(channelName, fabsdk.WithUser(user))
//	//cli, err := channel.New(ctx)
//	//if err != nil {
//	//	return channel.Response{}, err
//	//}
//	// 对区块链账本查询的操作（调用了链码的invoke），只返回结果
//	resp, err := t.Client.Query(channel.Request{
//		ChaincodeID: CC_Name,
//		Fcn:         fcn,
//		Args:        args,
//	})
//	if err != nil {
//		return channel.Response{}, err
//	}
//	//返回链码执行后的结果
//	return resp, nil
//}
//
////// 初始化服务
////func InitService(chaincode, channelID string, org *OrgInfo, sdk *fabsdk.FabricSDK) (*ServiceSetup, error) {
////	service := &ServiceSetup{
////		ChaincodeID: chaincode,
////	}
////	clientChnnelContext := sdk.ChannelContext(channelID, fabsdk.WithUser(org.OrgUser), fabsdk.WithOrg(org.OrgName))
////	client, err := channel.New(clientChnnelContext)
////	if err != nil {
////		return nil, fmt.Errorf(">> Failed to create new channel client: %s !!!!!!", err)
////	}
////	service.Client = client
////	return service, nil
////}
