#   About jd-sdk
京东开放平台golang版本sdk

### 京东api地址
<https://open.jd.com/#/doc/api?apiCateId=106&apiId=6146&apiName=jingdong.getUserEntryWechatGroupTrack>


##  demo

    var (
        AppKey      string = "Your AppKey"
        AppSecret   string = "Your AppSecret"
        AccessToken string = "Your accessToken"
    )

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

### ps

- 目前仅只支持了几个api，其他的需要自己手动接入，仿照写即可。

## License

Apache License, Version 2.0