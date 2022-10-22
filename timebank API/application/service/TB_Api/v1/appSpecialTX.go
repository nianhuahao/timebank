package v1

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"net/http"
	"strconv"
	"timebank/application/service/TB_Sdk"
	"timebank/application/service/TB_Utils"
)

//TransferAsset 转移或继承工分
type TransferSystem struct {
	//TransferID   string  //转移ID
	FromAssetId string `json:"fromAssetId"` //工分来源ID
	ToAssetId   string `json:"toAssetId"`   //工分去向ID
	AssetValue  string `json:"assetValue"`  //转移工分价值
	//TransferTime string  //转移工分时间
}

//RechargeSystem 充值服务
type RechargeSystem struct {
	//RechargeID    string  //充值ID
	ToUserID      string `json:"toUserID"`      //待充值用户的ID
	RechargeValue string `json:"rechargeValue"` //待充入工分总额
	//RechargeTime  string  //充值时间
}

func (t *ServiceSetup) TransferAsset(c *gin.Context) {
	respG := TB_Utils.Gin{C: c}
	body := new(TransferSystem)
	if err := c.ShouldBind(body); err != nil {
		respG.Response(http.StatusBadRequest, "Failed", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	a, _ := strconv.ParseFloat(body.AssetValue, 64)
	if body.FromAssetId == "" || body.ToAssetId == "" || a <= 0 {
		respG.Response(http.StatusBadRequest, "Failed", "请输入正确的参数")
		return
	}
	var bodyBytes [][]byte
	bodyBytes = append(bodyBytes, []byte(body.FromAssetId))
	bodyBytes = append(bodyBytes, []byte(body.ToAssetId))
	bodyBytes = append(bodyBytes, []byte(strconv.FormatFloat(a, 'E', -1, 64)))

	resp, err := t.Client.Execute(channel.Request{
		ChaincodeID: TB_Sdk.CC_Name,
		Fcn:         "TransferAsset",
		Args:        bodyBytes,
	})
	if err != nil {
		respG.Response(http.StatusBadRequest, "Failed", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}

	data := make(map[string]interface{}, 0)
	if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
		respG.Response(http.StatusInternalServerError, "Failed", err.Error())
		return
	}

	respG.Response(http.StatusOK, "Success", data)
}

func (t *ServiceSetup) InheritAsset(c *gin.Context) {
	respG := TB_Utils.Gin{C: c}
	body := new(TransferSystem)
	if err := c.ShouldBind(body); err != nil {
		respG.Response(http.StatusBadRequest, "Failed", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}

	if body.FromAssetId == "" || body.ToAssetId == "" {
		respG.Response(http.StatusBadRequest, "Failed", "请输入正确的参数")
		return
	}
	var bodyBytes [][]byte
	bodyBytes = append(bodyBytes, []byte(body.FromAssetId))
	bodyBytes = append(bodyBytes, []byte(body.ToAssetId))

	resp, err := t.Client.Execute(channel.Request{
		ChaincodeID: TB_Sdk.CC_Name,
		Fcn:         "InheritAsset",
		Args:        bodyBytes,
	})
	if err != nil {
		respG.Response(http.StatusBadRequest, "Failed", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}

	data := make(map[string]interface{}, 0)
	//var data []map[string]interface{}
	if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
		respG.Response(http.StatusInternalServerError, "Failed", err.Error())
		return
	}

	respG.Response(http.StatusOK, "Success", data)
}

func (t *ServiceSetup) RechargeAsset(c *gin.Context) {
	respG := TB_Utils.Gin{C: c}
	body := new(RechargeSystem)
	if err := c.ShouldBind(body); err != nil {
		respG.Response(http.StatusBadRequest, "Failed", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	a, _ := strconv.ParseFloat(body.RechargeValue, 64)
	if body.ToUserID == "" || a <= 0 {
		respG.Response(http.StatusBadRequest, "Failed", "请输入正确的KEY")
		return
	}
	var bodyBytes [][]byte
	bodyBytes = append(bodyBytes, []byte(body.ToUserID))
	bodyBytes = append(bodyBytes, []byte(strconv.FormatFloat(a, 'E', -1, 64)))

	resp, err := t.Client.Execute(channel.Request{
		ChaincodeID: TB_Sdk.CC_Name,
		Fcn:         "RechargeAsset",
		Args:        bodyBytes,
	})
	if err != nil {
		respG.Response(http.StatusBadRequest, "Failed", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}

	data := make(map[string]interface{}, 0)
	//data:=make(map[string]interface{},0)
	if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
		respG.Response(http.StatusInternalServerError, "Failed", err.Error())
		return
	}

	respG.Response(http.StatusOK, "Success", data)
}
