package TB_Sdk

import (
	"os"
)

const (
	CC_Name    = "timebank"
	CC_Version = "1.0.0"
)

var Orgs []*OrgInfo = []*OrgInfo{
	{
		OrgAdminUser:  "Admin",
		OrgName:       "Qsh",
		OrgMspId:      "QshMSP",
		OrgUser:       "User1",
		OrgPeerNum:    1,
		OrgAnchorFile: os.Getenv("GOPATH") + "/src/timebank/fixtures/channel-artifacts/QshAnchor.tx",
	},
	{
		OrgAdminUser:  "Admin",
		OrgName:       "Yh",
		OrgMspId:      "YhMSP",
		OrgUser:       "User1",
		OrgPeerNum:    1,
		OrgAnchorFile: os.Getenv("GOPATH") + "/src/timebank/fixtures/channel-artifacts/YhAnchor.tx",
	},
	{
		OrgAdminUser:  "Admin",
		OrgName:       "Gqc",
		OrgMspId:      "GqcMSP",
		OrgUser:       "User1",
		OrgPeerNum:    1,
		OrgAnchorFile: os.Getenv("GOPATH") + "/src/timebank/fixtures/channel-artifacts/GqcAnchor.tx",
	},
}

var Info SdkEnvInfo = SdkEnvInfo{
	ChannelID:        "mychannel",
	ChannelConfig:    os.Getenv("GOPATH") + "/src/timebank/fixtures/channel-artifacts/mychannel.tx",
	Orgs:             Orgs,
	OrdererAdminUser: "Admin",
	OrdererOrgName:   "OrdererOrg",
	OrdererEndpoint:  "orderer.jxnu.com",
	ChaincodeID:      CC_Name,
	ChaincodePath:    os.Getenv("GOPATH") + "/src/timebank/chaincode/",
	ChaincodeVersion: CC_Version,
}
