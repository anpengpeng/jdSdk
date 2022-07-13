package test

import (
	"fmt"
	"github.com/anpengpeng/jdSdk"
	jdUnionOpenGoodsQuery_request "github.com/anpengpeng/jdSdk/jdUnionOpenGoodsQuery/request"
	"testing"
)

const (
	AppKey      string = "Your AppSecret"
	AppSecret   string = "Your AppSecret"
	AccessToken string = "Your AccessToken"
)

func TestJdQuery(t *testing.T) {
	config := &jdSdk.JdBaseConfig{
		AppKey:         AppKey,
		AppSecret:      AppSecret,
		ApiVersion:     "1.0",
		AccessToken:    AccessToken,
		BaseUrl:        "https://api.jd.com/routerjson",
		ConnectTimeout: 2000,
	}
	queryRequest := jdUnionOpenGoodsQuery_request.New(config)
	params := queryRequest.GetParams()

	goodsQueryParam := jdUnionOpenGoodsQuery_request.GoodsReqDto{
		Keyword: "牛奶",
	}
	params.GoodsReqDto = goodsQueryParam
	execute, err := queryRequest.Execute()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(execute)
}
