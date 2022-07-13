package jdSdk

import (
	"encoding/json"
	"github.com/anpengpeng/jdSdk/utils"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type JdClient struct {
}

func NewJdClient() *JdClient {
	return &JdClient{}
}

func (j *JdClient) Request(params JdBaseApiRequest, isNeedAuth bool) (string, error) {
	timeStr := utils.GetTimestamp()
	buyParamJson, _ := json.Marshal(params.GetParamsObject())
	apiVersion := params.GetConfig().ApiVersion
	appKey := params.GetConfig().AppKey
	appSecret := params.GetConfig().AppSecret
	accessTokenStr := params.GetConfig().AccessToken
	method := params.GetMethodName()
	baseUrl := params.GetConfig().BaseUrl
	timeout := params.GetConfig().ConnectTimeout

	fields := map[string]string{
		"timestamp":         timeStr,
		"v":                 apiVersion,
		"sign_method":       "md5",
		"format":            "json",
		"method":            method,
		"360buy_param_json": string(buyParamJson),
		"app_key":           appKey,
	}
	accessToken := ""
	if isNeedAuth {
		fields["access_token"] = accessTokenStr
		accessToken = accessTokenStr
	}
	sign := utils.Sign(string(buyParamJson), accessToken, appKey, "json", params.GetMethodName(), "md5", timeStr, apiVersion, appSecret)
	fields["sign"] = sign

	return HttpPost(baseUrl, fields, timeout)
}

func HttpPost(urls string, data map[string]string, timeout int64) (string, error) {
	do := &http.Client{
		Timeout: time.Duration(timeout) * time.Millisecond,
	}

	var clusterInfo = url.Values{}
	for k, v := range data {
		clusterInfo.Add(k, v)
	}
	params := clusterInfo.Encode()
	resp, err := do.Post(urls, "application/x-www-form-urlencoded",
		strings.NewReader(params))
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

var DefaultJdApiClient *JdClient = NewJdClient()
