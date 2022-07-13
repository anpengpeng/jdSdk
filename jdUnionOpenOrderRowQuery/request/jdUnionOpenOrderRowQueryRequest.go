package request

import (
	"encoding/json"
	"errors"
	"github.com/anpengpeng/jdSdk"
	"github.com/anpengpeng/jdSdk/jdUnionOpenOrderRowQuery/response"
)

type UnionOpenOrderRowQueryRequest struct {
	jdSdk.BaseJdApiRequest
	Param *UnionOpenOrderRowQueryParam
}

func (j *UnionOpenOrderRowQueryRequest) GetMethodName() string {
	return "jd.union.open.order.row.query"
}

func New(config *jdSdk.JdBaseConfig) *UnionOpenOrderRowQueryRequest {
	request := UnionOpenOrderRowQueryRequest{
		Param: &UnionOpenOrderRowQueryParam{},
	}
	request.SetConfig(config)
	request.SetClient(jdSdk.DefaultJdApiClient)

	return &request
}

func (j *UnionOpenOrderRowQueryRequest) Execute() (result *response.UnionOpenOrderRowQueryResult, err error) {
	responseJson, err := j.GetClient().Request(j, false)
	if err != nil {
		return
	}

	resp := &response.UnionOpenOrderRowQueryResponseTopLevel{}
	if err = json.Unmarshal([]byte(responseJson), resp); err != nil {
		return
	}
	if resp.UnionOpenOrderRowQueryResponse.Result != "" {
		if err = json.Unmarshal([]byte(resp.UnionOpenOrderRowQueryResponse.Result), &result); err != nil {
			return
		}
		return
	} else {
		err = errors.New("result is null")
	}
	return

}

func (j *UnionOpenOrderRowQueryRequest) GetParamsObject() interface{} {
	return j.Param
}

func (j *UnionOpenOrderRowQueryRequest) GetParams() *UnionOpenOrderRowQueryParam {
	return j.Param
}

type UnionOpenOrderRowQueryParam struct {
	OrderReq OrderRowReq `json:"orderReq"`
}

type OrderRowReq struct {
	PageIndex    int    `json:"pageIndex"`              //页码
	PageSize     int    `json:"pageSize"`               //每页包含条数，上限为500
	Type         int    `json:"type"`                   //订单时间查询类型(1：下单时间，2：完成时间（购买用户确认收货时间），3：更新时间
	StartTime    string `json:"startTime"`              //开始时间 格式yyyy-MM-dd HH:mm:ss，与endTime间隔不超过1小时
	EndTime      string `json:"endTime"`                //结束时间 格式yyyy-MM-dd HH:mm:ss，与startTime间隔不超过1小时
	ChildUnionId int    `json:"childUnionId,omitempty"` //子推客unionID，传入该值可查询子推客的订单，注意不可和key同时传入。（需联系运营开通PID权限才能拿到数据）
	Key          string `json:"key,omitempty"`          //工具商传入推客的授权key，可帮助该推客查询订单，注意不可和childUnionid同时传入。（需联系运营开通工具商权限才能拿到数据）
	Fields       string `json:"fields,omitempty"`       //支持出参数据筛选，逗号','分隔，目前可用：goodsInfo（商品信息）,categoryInfo(类目信息）
	OrderId      int    `json:"orderId,omitempty"`      //订单号
}
