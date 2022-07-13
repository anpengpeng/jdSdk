package request

import (
	"encoding/json"
	"github.com/anpengpeng/jdSdk"
	"github.com/anpengpeng/jdSdk/jdUnionOpenPromotionBySubUnionIdGet/response"
)

type UnionOpenPromotionBySubUnionIdGetRequest struct {
	jdSdk.BaseJdApiRequest
	Param *UnionOpenPromotionBySubUnionIdGetParam
}

func (j *UnionOpenPromotionBySubUnionIdGetRequest) GetMethodName() string {
	return "jd.union.open.promotion.bysubunionid.get"
}

func New(config *jdSdk.JdBaseConfig) *UnionOpenPromotionBySubUnionIdGetRequest {
	request := UnionOpenPromotionBySubUnionIdGetRequest{
		Param: &UnionOpenPromotionBySubUnionIdGetParam{},
	}
	request.SetConfig(config)
	request.SetClient(jdSdk.DefaultJdApiClient)

	return &request
}

func (j *UnionOpenPromotionBySubUnionIdGetRequest) Execute() (*response.UnionOpenPromotionBySubUnionIdGetResponseResult, error) {
	responseJson, err := j.GetClient().Request(j, false)
	if err != nil {
		return nil, err
	}
	NewResponseObj := &response.UnionOpenPromotionBySubUnionIdGetResponseResult{}
	responseObj := map[string]response.UnionOpenPromotionBySubUnionIdGetResponse{}
	_ = json.Unmarshal([]byte(responseJson), &responseObj)
	for _, v := range responseObj {
		if v.GetResult != "" {
			_ = json.Unmarshal([]byte(v.GetResult), &NewResponseObj)
		}
		if v.QueryResult != "" {
			_ = json.Unmarshal([]byte(v.QueryResult), &NewResponseObj)
		}
		break
	}

	return NewResponseObj, err
}

func (j *UnionOpenPromotionBySubUnionIdGetRequest) GetParamsObject() interface{} {
	return j.Param
}

func (j *UnionOpenPromotionBySubUnionIdGetRequest) GetParams() *UnionOpenPromotionBySubUnionIdGetParam {
	return j.Param
}

type UnionOpenPromotionBySubUnionIdGetParam struct {
	PromotionCodeReq PromotionCodeReq `json:"promotionCodeReq"`
}

type PromotionCodeReq struct {
	MaterialId string `json:"materialId"` //推广物料url，例如活动链接、商品链接等；不支持仅传入skuid
	SubUnionId string `json:"subUnionId"` //子渠道标识，仅支持传入字母、数字、下划线或中划线，最多80个字符（不可包含空格），该参数会在订单行查询接口中展示（需申请权限，申请方法请见https://union.jd.com/helpcenter/13246-13247-46301）
	PositionId string `json:"positionId"` //推广位ID
	CouponUrl  string `json:"couponUrl"`
}
