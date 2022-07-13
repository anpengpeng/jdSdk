package response

type UnionOpenPromotionBySubUnionIdGetResponse struct {
	Code        string `json:"code"`
	GetResult   string `json:"getResult"`
	QueryResult string `json:"queryResult"`
}

type UnionOpenPromotionBySubUnionIdGetResponseResult struct {
	Code      int               `json:"code"`
	Message   string            `json:"message"`
	RequestId string            `json:"requestId"`
	Data      PromotionCodeResp `json:"data"`
}

type PromotionCodeResp struct {
	ShortURL      string `json:"shortURL"`
	ClickUrl      string `json:"clickURL"`
	JCommend      string `json:"jCommand"`
	JShortCommand string `json:"jShortCommand"`
}
