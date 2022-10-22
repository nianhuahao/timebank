package v1

import (
//	"bytes"
//	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"net/http"
	"strconv"
	"timebank/application/service/TB_Sdk"
	"timebank/application/service/TB_Utils"
)

type UpdateUser struct {
	UserID             string `json:"userID"`             //用户id
	UserName           string `json:"userName"`           //用户姓名
	UserIdentification string `json:"userIdentification"` //用户身份证号
	Sex                string `json:"sex"`                //用户性别
	Birthday           string `json:"birthday"`           //出生日期
	Address            string `json:"address"`            //住址
	Postcode           string `json:"postcode"`           //邮政编码
	Ability            string `json:"ability"`            //可提供的服务
	//StarSign           int      //星级标志

	//根据个人对组织的贡献、技能、资历年限等评定
	//星级标志分为：1、一颗心；2、两颗星；3、三颗星；4、四颗星；5、五颗星
	//星越多，个人对组织的贡献越多、技能越强、资历年限越高等

	//UserAsset     float64  //用户工分资产
	//Comment       []string //顾客对用户服务的评价
	//RecommenderID string //推荐人ID
}

type UpdateJob struct {
	JobID               string  `json:"jobID"`               //工作ID
	JobName             string  `json:"jobName"`             //工作名称
	JobUnitCost         float64 `json:"jobUnitCost"`         //工作每一定价
	DetailedDescription string  `json:"detailedDescription"` //工作描述
}

func (t *ServiceSetup) UpdateUserInfo(c *gin.Context) {
	respG := TB_Utils.Gin{C: c}
	body := new(UpdateUser)

	if err := c.ShouldBind(body); err != nil {
		respG.Response(http.StatusBadRequest, "Failed", fmt.Sprintf("参数出错%s", err.Error()))
	}
	if body.UserID == "" || body.UserName == "" || body.Sex == "" || body.UserIdentification == "" || body.Birthday == "" || body.Address == "" || body.Postcode == "" || body.Ability == "" {
		respG.Response(http.StatusBadRequest, "Failed", "请输入正确的个人信息 !!!")
		return
	}
	var bodyBytes [][]byte
	bodyBytes = append(bodyBytes, []byte(body.UserID))
	bodyBytes = append(bodyBytes, []byte(body.UserName))
	bodyBytes = append(bodyBytes, []byte(body.UserIdentification))
	bodyBytes = append(bodyBytes, []byte(body.Sex))
	bodyBytes = append(bodyBytes, []byte(body.Birthday))
	bodyBytes = append(bodyBytes, []byte(body.Address))
	bodyBytes = append(bodyBytes, []byte(body.Postcode))
	bodyBytes = append(bodyBytes, []byte(body.Ability))
	//bodyBytes = append(bodyBytes, []byte(body.RecommenderID))

	resp, err := t.Client.Execute(channel.Request{
		ChaincodeID: TB_Sdk.CC_Name,
		Fcn:         "UpdateUserInfo",
		Args:        bodyBytes,
	})
	if err != nil {
		respG.Response(http.StatusInternalServerError, "Failed", err.Error())
		return
	}

//	data := make(map[string]interface{}, 0)
	//var data map[string]interface{}
//	if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
//		respG.Response(http.StatusInternalServerError, "Failed", err.Error())
//		return
//	}
	var data interface{}
	data = string(resp.Payload)
	respG.Response(http.StatusOK, "Success", data)
}

func (t *ServiceSetup) UpdateServiceInfo(c *gin.Context) {
	respG := TB_Utils.Gin{C: c}
	body := new(UpdateJob)

	if err := c.ShouldBind(body); err != nil {
		respG.Response(http.StatusBadRequest, "Failed", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	if body.JobID == "" || body.JobName == "" || body.JobUnitCost <= 0 || body.DetailedDescription == "" {
		respG.Response(http.StatusBadRequest, "Failed", "请管理员输入正确的工作信息")
		return
	}

	var bodyBytes [][]byte
	bodyBytes = append(bodyBytes, []byte(body.JobID))
	bodyBytes = append(bodyBytes, []byte(body.JobName))
	bodyBytes = append(bodyBytes, []byte(strconv.FormatFloat(body.JobUnitCost, 'E', -1, 64)))
	bodyBytes = append(bodyBytes, []byte(body.DetailedDescription))

	resp, err := t.Client.Execute(channel.Request{
		ChaincodeID: TB_Sdk.CC_Name,
		Fcn:         "UpdateServiceInfo",
		Args:        bodyBytes,
	})
	if err != nil {
		respG.Response(http.StatusInternalServerError, "Failed", err.Error())
		return
	}
//	data := make(map[string]interface{}, 0)
	//var data map[string]interface{}
//	if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
//		respG.Response(http.StatusInternalServerError, "Failed", err.Error())
//		return
//	}
	var data interface{}
	data = string(resp.Payload)

	respG.Response(http.StatusOK, "Success", data)
}
