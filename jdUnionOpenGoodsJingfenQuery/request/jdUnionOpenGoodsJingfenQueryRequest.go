package request

import (
	"encoding/json"
	"github.com/anpengpeng/jdSdk"
	"github.com/anpengpeng/jdSdk/jdUnionOpenGoodsJingfenQuery/response"
)

type UnionOpenGoodsJingfenQueryRequest struct {
	jdSdk.BaseJdApiRequest
	Param *UnionOpenGoodsJingfenQueryParam
}

func (j *UnionOpenGoodsJingfenQueryRequest) GetMethodName() string {
	return "jd.union.open.goods.jingfen.query"
}

func New(config *jdSdk.JdBaseConfig) *UnionOpenGoodsJingfenQueryRequest {
	request := UnionOpenGoodsJingfenQueryRequest{
		Param: &UnionOpenGoodsJingfenQueryParam{},
	}
	request.SetConfig(config)
	request.SetClient(jdSdk.DefaultJdApiClient)

	return &request
}

func (j *UnionOpenGoodsJingfenQueryRequest) Execute() (*response.UnionOpenGoodsJingfenQueryResponseResult, error) {
	responseJson, err := j.GetClient().Request(j, false)
	if err != nil {
		return nil, err
	}
	NewResponseObj := &response.UnionOpenGoodsJingfenQueryResponseResult{}
	responseObj := map[string]response.UnionOpenGoodsJingfenQueryResponse{}
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

func (j *UnionOpenGoodsJingfenQueryRequest) GetParamsObject() interface{} {
	return j.Param
}

func (j *UnionOpenGoodsJingfenQueryRequest) GetParams() *UnionOpenGoodsJingfenQueryParam {
	return j.Param
}

type UnionOpenGoodsJingfenQueryParam struct {
	GoodsReq GoodsReq `json:"goodsReq"`
}

type GoodsReq struct {
	EliteId      int    `json:"eliteId"` //频道ID:1-好券商品,2-精选卖场,10-9.9包邮,15-京东配送,22-实时热销榜,23-为你推荐,24-数码家电,25-超市,26-母婴玩具,27-家具日用,28-美妆穿搭,30-图书文具,31-今日必推,32-京东好物,33-京东秒杀,34-拼购商品,40-高收益榜,41-自营热卖榜,108-秒杀进行中,109-新品首发,110-自营,112-京东爆品,125-首购商品,129-高佣榜单,130-视频商品,153-历史最低价商品榜,210-极速版商品,238-新人价商品,247-京喜9.9,249-京喜秒杀,315-秒杀未开始,340-时尚趋势品,341-3C新品,342-智能新品,343-3C长尾商品,345-时尚新品,346-时尚爆品,426-京喜自营,1001-选品库
	PageIndex    int    `json:"pageIndex"`
	PageSize     int    `json:"pageSize"`     //每页数量，默认20，上限50，建议20
	SortName     string `json:"sortName"`     //排序字段(price：单价, commissionShare：佣金比例, commission：佣金， inOrderCount30Days：30天引单量， inOrderComm30Days：30天支出佣金)
	Sort         string `json:"sort"`         //asc,desc升降序,默认降序
	Fields       string `json:"fields"`       //支持出参数据筛选，逗号','分隔，目前可用：videoInfo(视频信息),hotWords(热词),similar(相似推荐商品),documentInfo(段子信息)，skuLabelInfo（商品标签），promotionLabelInfo（商品促销标签）
	ForbidTypes  string `json:"forbidTypes"`  //10微信京东购物小程序禁售，11微信京喜小程序禁售
	GroupId      int    `json:"groupId"`      //选品库id（仅对eliteId=1001有效，且必传）
	OwnerUnionId int    `json:"ownerUnionId"` //groupId创建者的UnionId
	Pid          string `json:"pid"`
}
