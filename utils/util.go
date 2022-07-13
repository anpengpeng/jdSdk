package utils

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"strings"
	"time"
)

func GetTimestamp() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func Sign(paramJson, accessToken, appKey, format, method, signMethod, timestamp, v, appSecret string) string {
	signStr := appSecret + "360buy_param_json" + paramJson
	if accessToken != "" {
		signStr += "access_token" + accessToken
	}
	signStr += "app_key" + appKey + "format" + format + "method" + method + "sign_method" + signMethod + "timestamp" + timestamp + "v" + v + appSecret

	s := Md5(signStr)
	return strings.ToUpper(s)
}

func Md5(s string) string {
	h := md5.New()
	_, _ = io.WriteString(h, s)
	return hex.EncodeToString(h.Sum(nil))
}
