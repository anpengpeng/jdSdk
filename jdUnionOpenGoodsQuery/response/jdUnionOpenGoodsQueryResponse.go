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
	CurrentPrice     int `json:"currentPrice"`     //预售价格
	Earnest          int `json:"earnest"`          //订金金额（定金不能超过预售总价的20%）
	PreSalePayType   int `json:"preSalePayType"`   //预售支付类型：1.仅全款 2.定金、全款均可 5.一阶梯仅定金
	DiscountType     int `json:"discountType"`     //1: 定金膨胀 2: 定金立减
	DepositWorth     int `json:"depositWorth"`     //定金膨胀金额（定金可抵XXX）【废弃】
	PreAmountDeposit int `json:"preAmountDeposit"` // 立减金额
	PreSaleStartTime int `json:"preSaleStartTime"` //定金开始时间
	PreSaleEndTime   int `json:"preSaleEndTime"`   //定金结束时间
	BalanceStartTime int `json:"balanceStartTime"` // 尾款开始时间
	BalanceEndTime   int `json:"balanceEndTime"`   // 尾款结束时间
	ShipTime         int `json:"shipTime"`         // 预计发货时间
	PreSaleStatus    int `json:"preSaleStatus"`    // 预售状态（0 未开始；1 预售中；2 预售结束；3 尾款进行中；4 尾款结束）
	AmountDeposit    int `json:"amountDeposit"`    // 定金膨胀金额（定金可抵XXX）
}
