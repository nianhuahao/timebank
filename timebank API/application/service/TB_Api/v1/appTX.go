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

//Servicing 进行服务
type Servicing struct {
	ServicingType    string `json:"servicingType"`    //服务类型
	ServicingOlderID string `json:"servicingOlderID"` //发起服务老人ID
	ServicingVolID   string `json:"servicingVolID"`   //接受服务志愿者ID
	ServicingCount   string `json:"servicingCount"`   //服务计数
	//StartTime        string  //发起服务开始时间
	ServicingValue string `json:"servicingValue"` //服务所需工分
	//ServicingState   string  //服务状态
}

//ServiceTrade 交易，服务与被服务的工分交易
type ServiceTrade struct {
	//TxID      string //交易ID
	ServeID string `json:"serveID"` //服务人员的ID
	CustID  string `json:"custID"`  //顾客ID
	TxType  string `json:"txType"`  //服务类型
	//WorkHours int    //服务计数

	//如果是按时计工分，即为工作时长
	//如果是按件计工分，即为总件数
	//如果是按次计工分，即为总次数

	//EndTime      string  //服务结束时间
	//WorkValue    float64 //服务所需工分
	ServeComment string `json:"serveComment"` //顾客评价
}

func (t *ServiceSetup) CreateServicing(c *gin.Context) {
	respG := TB_Utils.Gin{C: c}
	body := new(Servicing)
	if err := c.ShouldBind(body); err != nil {
		respG.Response(http.StatusBadRequest, "Failed", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	a, _ := strconv.Atoi(body.ServicingCount)
	if body.ServicingOlderID == "" || body.ServicingType == "" || a <= 0 {
		respG.Response(http.StatusBadRequest, "Failed", "请输入正确的参数")
		return
	}
	var bodyBytes [][]byte
	bodyBytes = append(bodyBytes, []byte(body.ServicingType))
	bodyBytes = append(bodyBytes, []byte(body.ServicingOlderID))
	bodyBytes = append(bodyBytes, []byte(strconv.FormatInt(int64(a), 10)))

	resp, err := t.Client.Execute(channel.Request{
		ChaincodeID: TB_Sdk.CC_Name,
		Fcn:         "CreateServicing",
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

func (t *ServiceSetup) AcceptServicing(c *gin.Context) {
	respG := TB_Utils.Gin{C: c}
	body := new(Servicing)
	if err := c.ShouldBind(body); err != nil {
		respG.Response(http.StatusBadRequest, "Failed", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}

	if body.ServicingOlderID == "" || body.ServicingVolID == "" || body.ServicingType == "" {
		respG.Response(http.StatusBadRequest, "Failed", "请输入正确的参数")
		return
	}
	var bodyBytes [][]byte
	bodyBytes = append(bodyBytes, []byte(body.ServicingType))
	bodyBytes = append(bodyBytes, []byte(body.ServicingVolID))
	bodyBytes = append(bodyBytes, []byte(body.ServicingOlderID))

	resp, err := t.Client.Execute(channel.Request{
		ChaincodeID: TB_Sdk.CC_Name,
		Fcn:         "AcceptServicing",
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

func (t *ServiceSetup) DoneServicing(c *gin.Context) {
	respG := TB_Utils.Gin{C: c}
	body := new(ServiceTrade)
	if err := c.ShouldBind(body); err != nil {
		respG.Response(http.StatusBadRequest, "Failed", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}

	if body.CustID == "" || body.ServeComment == "" || body.ServeID == "" || body.TxType == "" {
		respG.Response(http.StatusBadRequest, "Failed", "请输入正确的参数")
		return
	}
	var bodyBytes [][]byte
	bodyBytes = append(bodyBytes, []byte(body.TxType))
	bodyBytes = append(bodyBytes, []byte(body.CustID))
	bodyBytes = append(bodyBytes, []byte(body.ServeID))
	bodyBytes = append(bodyBytes, []byte(body.ServeComment))

	resp, err := t.Client.Execute(channel.Request{
		ChaincodeID: TB_Sdk.CC_Name,
		Fcn:         "DoneServicing",
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

func (t *ServiceSetup) CloseServicing(c *gin.Context) {
	respG := TB_Utils.Gin{C: c}
	body := new(Servicing)
	if err := c.ShouldBind(body); err != nil {
		respG.Response(http.StatusBadRequest, "Failed", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}

	if body.ServicingOlderID == "" || body.ServicingType == "" {
		respG.Response(http.StatusBadRequest, "Failed", "请输入正确的参数")
		return
	}
	var bodyBytes [][]byte
	bodyBytes = append(bodyBytes, []byte(body.ServicingType))
	bodyBytes = append(bodyBytes, []byte(body.ServicingOlderID))

	resp, err := t.Client.Execute(channel.Request{
		ChaincodeID: TB_Sdk.CC_Name,
		Fcn:         "CloseServicing",
		Args:        bodyBytes,
	})
	if err != nil {
		respG.Response(http.StatusBadRequest, "Failed", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}

	//data := make(map[string]interface{}, 0)
	//var data []map[string]interface{}
	//if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
	//	respG.Response(http.StatusInternalServerError, "Failed", err.Error())
	//	return
	//}

	var data interface{}
	data = string(resp.Payload)
	respG.Response(http.StatusOK, "Success", data)
}
