package response

import "github.com/anpengpeng/jdSdk/jdUnionOpenGoodsQuery/response"

type UnionOpenGoodsJingfenQueryResponse struct {
	Code        string `json:"code"`
	GetResult   string `json:"getResult"`
	QueryResult string `json:"queryResult"`
}

type UnionOpenGoodsJingfenQueryResponseResult struct {
	Code      int                              `json:"code"`
	Message   string                           `json:"message"`
	RequestId string                           `json:"requestId"`
	Data      []UnionOpenGoodsJingfenQueryResp `json:"data"`
}

type UnionOpenGoodsJingfenQueryResp struct {
	BrandCode    string `json:"brandCode"`
	BrandName    string `json:"brandName"`
	CategoryInfo struct {
		Cid1     int    `json:"cid1"`
		Cid1Name string `json:"cid1Name"`
		Cid2     int    `json:"cid2"`
		Cid2Name string `json:"cid2Name"`
		Cid3     int    `json:"cid3"`
		Cid3Name string `json:"cid3Name"`
	} `json:"categoryInfo"`
	Comments       int `json:"comments"`
	CommissionInfo struct {
		Commission          float64 `json:"commission"`
		CommissionShare     float64 `json:"commissionShare"`
		CouponCommission    float64 `json:"couponCommission"`
		EndTime             int64   `json:"endTime"`
		IsLock              int     `json:"isLock"`
		PlusCommissionShare float64 `json:"plusCommissionShare"`
		StartTime           int64   `json:"startTime"`
	} `json:"commissionInfo"`
	CouponInfo struct {
		CouponList []response.JdCoupon `json:"couponList"`
	} `json:"couponInfo"`
	DeliveryType      int     `json:"deliveryType"`
	EliteType         []int   `json:"eliteType,omitempty"`
	ForbidTypes       []int   `json:"forbidTypes"`
	GoodCommentsShare float64 `json:"goodCommentsShare"`
	ImageInfo         struct {
		ImageList []struct {
			Url string `json:"url"`
		} `json:"imageList"`
		WhiteImage string `json:"whiteImage,omitempty"`
	} `json:"imageInfo"`
	InOrderComm30Days  float64             `json:"inOrderComm30Days"`
	InOrderCount30Days int                 `json:"inOrderCount30Days"`
	IsHot              int                 `json:"isHot"`
	IsJdSale           int                 `json:"isJdSale"`
	MaterialUrl        string              `json:"materialUrl"`
	Owner              string              `json:"owner"`
	PinGouInfo         response.PinGouInfo `json:"pinGouInfo"`
	PriceInfo          struct {
		HistoryPriceDay   int     `json:"historyPriceDay,omitempty"`
		LowestCouponPrice float64 `json:"lowestCouponPrice"`
		LowestPrice       float64 `json:"lowestPrice"`
		LowestPriceType   int     `json:"lowestPriceType"`
		Price             float64 `json:"price"`
	} `json:"priceInfo"`
	ShopInfo struct {
		AfsFactorScoreRankGrade       string  `json:"afsFactorScoreRankGrade,omitempty"`
		AfterServiceScore             string  `json:"afterServiceScore,omitempty"`
		CommentFactorScoreRankGrade   string  `json:"commentFactorScoreRankGrade,omitempty"`
		LogisticsFactorScoreRankGrade string  `json:"logisticsFactorScoreRankGrade,omitempty"`
		LogisticsLvyueScore           string  `json:"logisticsLvyueScore,omitempty"`
		ScoreRankRate                 string  `json:"scoreRankRate,omitempty"`
		ShopId                        int     `json:"shopId"`
		ShopLabel                     string  `json:"shopLabel"`
		ShopLevel                     float64 `json:"shopLevel"`
		ShopName                      string  `json:"shopName"`
		UserEvaluateScore             string  `json:"userEvaluateScore,omitempty"`
	} `json:"shopInfo"`
	SkuId     int64  `json:"skuId"`
	SkuName   string `json:"skuName"`
	Spuid     int64  `json:"spuid"`
	VideoInfo struct {
	} `json:"videoInfo"`
	SecondPriceInfoList []struct {
		SecondPrice     float64 `json:"secondPrice"`
		SecondPriceType int     `json:"secondPriceType"`
	} `json:"secondPriceInfoList,omitempty"`
	SeckillInfo struct {
		SeckillEndTime   int64   `json:"seckillEndTime"`
		SeckillOriPrice  float64 `json:"seckillOriPrice"`
		SeckillPrice     float64 `json:"seckillPrice"`
		SeckillStartTime int64   `json:"seckillStartTime"`
	} `json:"seckillInfo,omitempty"`
}
