package umeng

import (
	"fmt"
	"crypto/md5"
	"encoding/hex"
	"github.com/smartwalle/going/http"
	"strings"
)

const (
	UMENG_MSG_SEND_API_URL = "http://msg.umeng.com/api/send"
)

func PushUnicastMessage(appKey, appSecret, deviceToken string, productionMode bool, payload map[string]interface{}) (map[string]interface{}) {
	var message = NewUMengMessage(appKey, appSecret, UMENG_MESSAGE_TYPE_UNICAST, productionMode)
	message.DeviceTokens = deviceToken
	message.Payload = payload
	return PushMessage(message)
}

func PushListcastMessage(appKey, appSecret string, deviceTokens []string, productionMode bool, payload map[string]interface{}) (map[string]interface{}) {
	var message = NewUMengMessage(appKey, appSecret, UMENG_MESSAGE_TYPE_LISTCAST, productionMode)
	message.DeviceTokens = strings.Join(deviceTokens, ",")
	message.Payload = payload
	return PushMessage(message)
}

func PushBroadcastMessage(appKey, appSecret string, productionMode bool, payload map[string]interface{}) (map[string]interface{}) {
	var message = NewUMengMessage(appKey, appSecret, UMENG_MESSAGE_TYPE_BROADCAST, productionMode)
	message.Payload = payload
	return PushMessage(message)
}

func PushMessage(message *UMengMessage) (map[string]interface{}) {
	var sign = sign("POST", UMENG_MSG_SEND_API_URL, message.JSON(), message.AppSecret)
	var client = http.NewClient()
	client.SetMethod("POST")
	client.SetURLString(UMENG_MSG_SEND_API_URL+"?sign="+sign)
	client.SetBody(message.JSON())

	var results, _ = client.DoJsonRequest()
	return results
}

func sign(method, url, postBody, appSecret string) string {
	var sign = fmt.Sprintf("%s%s%s%s", method, url, postBody, appSecret)
	var m = md5.New()
	m.Write([]byte(sign))
	sign = hex.EncodeToString(m.Sum(nil))
	return sign
}