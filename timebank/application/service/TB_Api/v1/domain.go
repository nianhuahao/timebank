package v1

import (
	"fmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"timebank/application/service/TB_Sdk"
)

type ServiceSetup struct {
	ChaincodeID string
	Client      *channel.Client
}

//var t *ServiceSetup

// 初始化服务
func InitService(chaincode, channelID string, org *TB_Sdk.OrgInfo, sdk *fabsdk.FabricSDK) (*ServiceSetup, error) {
	service := &ServiceSetup{
		ChaincodeID: chaincode,
	}
	clientChnnelContext := sdk.ChannelContext(channelID, fabsdk.WithUser(org.OrgUser), fabsdk.WithOrg(org.OrgName))
	client, err := channel.New(clientChnnelContext)
	if err != nil {
		return nil, fmt.Errorf(">> Failed to create new channel client: %s !!!!!!", err)
	}
	service.Client = client
	return service, nil
}
