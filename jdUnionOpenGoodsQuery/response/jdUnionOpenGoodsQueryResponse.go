package response

type UnionOpenGoodsQueryResponse struct {
	Code        string `json:"code"`
	GetResult   string `json:"getResult"`
	QueryResult string `json:"queryResult"`
}

type UnionOpenGoodsQueryResponseResult struct {
	Code      int                       `json:"code"`
	Message   string                    `json:"message"`
	RequestId string                    `json:"requestId"`
	Data      []UnionOpenGoodsQueryResp `json:"data"`
}

type UnionOpenGoodsQueryResp struct {
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
		CouponList []JdCoupon `json:"couponList"`
	} `json:"couponInfo"`
	DeliveryType      int     `json:"deliveryType"`
	ForbidTypes       []int   `json:"forbidTypes"`
	GoodCommentsShare float64 `json:"goodCommentsShare"`
	ImageInfo         struct {
		ImageList []struct {
			Url string `json:"url"`
		} `json:"imageList"`
		WhiteImage string `json:"whiteImage,omitempty"`
	} `json:"imageInfo"`
	InOrderComm30Days  float64    `json:"inOrderComm30Days"`
	InOrderCount30Days int        `json:"inOrderCount30Days"`
	IsHot              int        `json:"isHot"`
	IsJdSale           int        `json:"isJdSale"`
	MaterialUrl        string     `json:"materialUrl"`
	Owner              string     `json:"owner"`
	PinGouInfo         PinGouInfo `json:"pinGouInfo"`
	PingGouInfo        struct {
	} `json:"pingGouInfo"`
	PriceInfo struct {
		LowestCouponPrice float64 `json:"lowestCouponPrice"`
		LowestPrice       float64 `json:"lowestPrice"`
		LowestPriceType   int     `json:"lowestPriceType"`
		Price             float64 `json:"price"`
		HistoryPriceDay   int     `json:"historyPriceDay,omitempty"`
	} `json:"priceInfo"`
	ShopInfo struct {
		ShopId                        int     `json:"shopId"`
		ShopLabel                     string  `json:"shopLabel"`
		ShopLevel                     float64 `json:"shopLevel"`
		ShopName                      string  `json:"shopName"`
		AfsFactorScoreRankGrade       string  `json:"afsFactorScoreRankGrade,omitempty"`
		AfterServiceScore             string  `json:"afterServiceScore,omitempty"`
		CommentFactorScoreRankGrade   string  `json:"commentFactorScoreRankGrade,omitempty"`
		LogisticsFactorScoreRankGrade string  `json:"logisticsFactorScoreRankGrade,omitempty"`
		LogisticsLvyueScore           string  `json:"logisticsLvyueScore,omitempty"`
		ScoreRankRate                 string  `json:"scoreRankRate,omitempty"`
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
	EliteType   []int       `json:"eliteType,omitempty"`
	PreSaleInfo PreSaleInfo `json:"preSaleInfo"`
}

type PinGouInfo struct {
	PinGouPrice     float64 `json:"pingouPrice"`
	PinGouTmCount   int     `json:"pingouTmCount"`
	PinGouUrl       string  `json:"pingouUrl"`
	PinGouStartTime int     `json:"pingouStartTime"`
	PinGouEndTime   int     `json:"pingouEndTime"`
}

type JdCoupon struct {
	BindType      int     `json:"bindType"`
	Discount      float64 `json:"discount"`
	GetEndTime    int64   `json:"getEndTime"`
	GetStartTime  int64   `json:"getStartTime"`
	HotValue      int     `json:"hotValue,omitempty"`
	IsBest        int     `json:"isBest"`
	IsInputCoupon int     `json:"isInputCoupon"`
	Link          string  `json:"link"`
	PlatformType  int     `json:"platformType"`
	Quota         float64 `json:"quota"`
	UseEndTime    int64   `json:"useEndTime"`
	UseStartTime  int64   `json:"useStartTime"`
}

type PreSaleInfo struct {
	CurrentPrice     int `json:"currentPrice"`     //????????????
	Earnest          int `json:"earnest"`          //????????????????????????????????????????????????20%???
	PreSalePayType   int `json:"preSalePayType"`   //?????????????????????1.????????? 2.????????????????????? 5.??????????????????
	DiscountType     int `json:"discountType"`     //1: ???????????? 2: ????????????
	DepositWorth     int `json:"depositWorth"`     //?????????????????????????????????XXX???????????????
	PreAmountDeposit int `json:"preAmountDeposit"` // ????????????
	PreSaleStartTime int `json:"preSaleStartTime"` //??????????????????
	PreSaleEndTime   int `json:"preSaleEndTime"`   //??????????????????
	BalanceStartTime int `json:"balanceStartTime"` // ??????????????????
	BalanceEndTime   int `json:"balanceEndTime"`   // ??????????????????
	ShipTime         int `json:"shipTime"`         // ??????????????????
	PreSaleStatus    int `json:"preSaleStatus"`    // ???????????????0 ????????????1 ????????????2 ???????????????3 ??????????????????4 ???????????????
	AmountDeposit    int `json:"amountDeposit"`    // ?????????????????????????????????XXX???
}
