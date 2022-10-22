package v1

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"net/http"
	"timebank/application/service/TB_Sdk"
	"timebank/application/service/TB_Utils"
)

type UserListQueryRequestBody struct {
	UserKey string `json:"userKey"`
}

type OrgListQueryRequestBody struct {
	OrgKey string `json:"orgKey"`
}

type ManagerListQueryRequestBody struct {
	ManagerKey string `json:"managerKey"`
}

type JobListQueryRequestBody struct {
	Job string `json:"job"`
}

type ServicingStatusListQueryRequestBody struct {
	ElderId string `json:"elderId"`
	JobId   string `json:"jobId"`
}

type TxListQueryRequestBody struct {
	TxId        string `json:"txId"`
	ElderId     string `json:"elderId"`
	VolunteerId string `json:"volunteerId"`
}

type SpecialTxListQueryRequestBody struct {
	SpecialTxId string `json:"specialTxId"`
}

type RechargeListQueryRequestBody struct {
	RechargeId string `json:"rechargeId"`
}

type GetUpdateHistoryListQueryRequestBody struct {
	UpdateThingName string `json:"updateThingName"`
	UpdateThingId   string `json:"updateThingId"`
}

func (t *ServiceSetup) UserList(c *gin.Context) {
	respG := TB_Utils.Gin{C: c}
	body := new(UserListQueryRequestBody)
	if err := c.ShouldBind(body); err != nil {
		respG.Response(http.StatusBadRequest, "Failed", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	if body.UserKey == "" {
		respG.Response(http.StatusBadRequest, "Failed", "请输入正确的KEY")
		return
	}

	var bodyBytes [][]byte
	if body.UserKey != "" {
		bodyBytes = append(bodyBytes, []byte(body.UserKey))
	}

	resp, err := t.Client.Query(channel.Request{
		ChaincodeID: TB_Sdk.CC_Name,
		Fcn:         "UserList",
		Args:        bodyBytes,
	})
	if err != nil {
		respG.Response(http.StatusBadRequest, "Failed", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}

	data := make([]map[string]interface{}, 0)
	//var data []map[string]interface{}
	if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
		respG.Response(http.StatusInternalServerError, "Failed", err.Error())
		return
	}

	respG.Response(http.StatusOK, "Success", data)
}

func (t *ServiceSetup) OrgList(c *gin.Context) {
	respG := TB_Utils.Gin{C: c}
	body := new(OrgListQueryRequestBody)
	if err := c.ShouldBind(body); err != nil {
		respG.Response(http.StatusBadRequest, "Failed", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	if body.OrgKey == "" {
		respG.Response(http.StatusBadRequest, "Failed", "请输入正确的KEY")
		return
	}

	var bodyBytes [][]byte
	bodyBytes = append(bodyBytes, []byte(body.OrgKey))

	resp, err := t.Client.Query(channel.Request{
		ChaincodeID: TB_Sdk.CC_Name,
		Fcn:         "OrgList",
		Args:        bodyBytes,
	})
	if err != nil {
		respG.Response(http.StatusBadRequest, "Failed", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}

	data := make([]map[string]interface{}, 0)
	//var data []map[string]interface{}
	if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
		respG.Response(http.StatusInternalServerError, "Failed", err.Error())
		return
	}

	respG.Response(http.StatusOK, "Success", data)
}

func (t *ServiceSetup) ManagerList(c *gin.Context) {
	respG := TB_Utils.Gin{C: c}
	body := new(ManagerListQueryRequestBody)
	if err := c.ShouldBind(body); err != nil {
		respG.Response(http.StatusBadRequest, "Failed", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	if body.ManagerKey == "" {
		respG.Response(http.StatusBadRequest, "Failed", "请输入正确的KEY")
		return
	}

	var bodyBytes [][]byte
	bodyBytes = append(bodyBytes, []byte(body.ManagerKey))

	resp, err := t.Client.Query(channel.Request{
		ChaincodeID: TB_Sdk.CC_Name,
		Fcn:         "ManagerList",
		Args:        bodyBytes,
	})
	if err != nil {
		respG.Response(http.StatusBadRequest, "Failed", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}

	data := make([]map[string]interface{}, 0)
	//var data []map[string]interface{}
	if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
		respG.Response(http.StatusInternalServerError, "Failed", err.Error())
		return
	}

	respG.Response(http.StatusOK, "Success", data)
}

func (t *ServiceSetup) JobList(c *gin.Context) {
	respG := TB_Utils.Gin{C: c}
	body := new(JobListQueryRequestBody)
	if err := c.ShouldBind(body); err != nil {
		respG.Response(http.StatusBadRequest, "Failed", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}

	var bodyBytes [][]byte
	if body.Job != "" {
		bodyBytes = append(bodyBytes, []byte(body.Job))
	}

	resp, err := t.Client.Query(channel.Request{
		ChaincodeID: TB_Sdk.CC_Name,
		Fcn:         "ServiceList",
		Args:        bodyBytes,
	})
	if err != nil {
		respG.Response(http.StatusBadRequest, "Failed", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}

	data := make([]map[string]interface{}, 0)
	//var data []map[string]interface{}
	if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
		respG.Response(http.StatusInternalServerError, "Failed", err.Error())
		return
	}

	respG.Response(http.StatusOK, "Success", data)
}

func (t *ServiceSetup) ServicingStatusList(c *gin.Context) {
	respG := TB_Utils.Gin{C: c}
	body := new(ServicingStatusListQueryRequestBody)
	if err := c.ShouldBind(body); err != nil {
		respG.Response(http.StatusBadRequest, "Failed", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}

	var bodyBytes [][]byte
	if body.ElderId != "" && body.JobId != "" {
		bodyBytes = append(bodyBytes, []byte(body.ElderId))
		bodyBytes = append(bodyBytes, []byte(body.JobId))
	}

	resp, err := t.Client.Query(channel.Request{
		ChaincodeID: TB_Sdk.CC_Name,
		Fcn:         "QueryServicingStatus",
		Args:        bodyBytes,
	})
	if err != nil {
		respG.Response(http.StatusBadRequest, "Failed", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}

	data := make([]map[string]interface{}, 0)
	//var data []map[string]interface{}
	if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
		respG.Response(http.StatusInternalServerError, "Failed", err.Error())
		return
	}

	respG.Response(http.StatusOK, "Success", data)
}

func (t *ServiceSetup) TxList(c *gin.Context) {
	respG := TB_Utils.Gin{C: c}
	body := new(TxListQueryRequestBody)
	if err := c.ShouldBind(body); err != nil {
		respG.Response(http.StatusBadRequest, "Failed", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}

	var bodyBytes [][]byte
	if body.TxId != "" && body.ElderId != "" && body.VolunteerId != "" {
		bodyBytes = append(bodyBytes, []byte(body.TxId))
		bodyBytes = append(bodyBytes, []byte(body.ElderId))
		bodyBytes = append(bodyBytes, []byte(body.VolunteerId))
	}

	resp, err := t.Client.Query(channel.Request{
		ChaincodeID: TB_Sdk.CC_Name,
		Fcn:         "QueryServiceTrade",
		Args:        bodyBytes,
	})
	if err != nil {
		respG.Response(http.StatusBadRequest, "Failed", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}

	data := make([]map[string]interface{}, 0)
	//var data []map[string]interface{}
	if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
		respG.Response(http.StatusInternalServerError, "Failed", err.Error())
		return
	}

	respG.Response(http.StatusOK, "Success", data)
}

func (t *ServiceSetup) SpecialTxList(c *gin.Context) {
	respG := TB_Utils.Gin{C: c}
	body := new(SpecialTxListQueryRequestBody)
	if err := c.ShouldBind(body); err != nil {
		respG.Response(http.StatusBadRequest, "Failed", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	if body.SpecialTxId == "" {
		respG.Response(http.StatusBadRequest, "Failed", "请输入正确的KEY")
		return
	}

	var bodyBytes [][]byte
	bodyBytes = append(bodyBytes, []byte(body.SpecialTxId))

	resp, err := t.Client.Query(channel.Request{
		ChaincodeID: TB_Sdk.CC_Name,
		Fcn:         "SpecialTradeList",
		Args:        bodyBytes,
	})
	if err != nil {
		respG.Response(http.StatusBadRequest, "Failed", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}

	data := make([]map[string]interface{}, 0)
	//var data []map[string]interface{}
	if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
		respG.Response(http.StatusInternalServerError, "Failed", err.Error())
		return
	}

	respG.Response(http.StatusOK, "Success", data)
}

func (t *ServiceSetup) RechargeList(c *gin.Context) {
	respG := TB_Utils.Gin{C: c}
	body := new(RechargeListQueryRequestBody)
	if err := c.ShouldBind(body); err != nil {
		respG.Response(http.StatusBadRequest, "Failed", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}

	var bodyBytes [][]byte
	if body.RechargeId != "" {
		bodyBytes = append(bodyBytes, []byte(body.RechargeId))
	}

	resp, err := t.Client.Query(channel.Request{
		ChaincodeID: TB_Sdk.CC_Name,
		Fcn:         "RechargeList",
		Args:        bodyBytes,
	})
	if err != nil {
		respG.Response(http.StatusBadRequest, "Failed", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}

	//var data interface{}
	data := make([]map[string]interface{}, 0)
//	var data []map[string]interface{}
	if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
		respG.Response(http.StatusInternalServerError, "Failed", err.Error())
		return
	}

	//data=string(resp.Payload)

	respG.Response(http.StatusOK, "Success", data)
}

func (t *ServiceSetup) GetUpdateHistoryList(c *gin.Context) {
	respG := TB_Utils.Gin{C: c}
	body := new(GetUpdateHistoryListQueryRequestBody)
	if err := c.ShouldBind(body); err != nil {
		respG.Response(http.StatusBadRequest, "Failed", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}

	if body.UpdateThingName == "" || body.UpdateThingId == "" {
		respG.Response(http.StatusBadRequest, "Failed", "请输入正确的KEY")
		return
	}
	var bodyBytes [][]byte
	bodyBytes = append(bodyBytes, []byte(body.UpdateThingName))
	bodyBytes = append(bodyBytes, []byte(body.UpdateThingId))

	resp, err := t.Client.Query(channel.Request{
		ChaincodeID: TB_Sdk.CC_Name,
		Fcn:         "GetUpdateHistory",
		Args:        bodyBytes,
	})
	if err != nil {
		respG.Response(http.StatusBadRequest, "Failed", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}

	//data := make(map[string]interface{}, 0)
//	var data []map[string]interface{}
	//if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
	//	respG.Response(http.StatusInternalServerError, "Failed", err.Error())
	//	return
	//}
	var data interface{}
	data=string(resp.Payload)

	respG.Response(http.StatusOK, "Success", data)
}
