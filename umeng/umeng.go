package umeng

import (
	"fmt"
	"strings"
	"crypto/md5"
	"encoding/hex"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"bytes"
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

func PushMessage(message *UMengMessage) (results map[string]interface{}) {
	if message == nil {
		return nil
	}

	b, err := json.Marshal(message)
	if err != nil {
		return nil
	}
	var buf = bytes.NewBuffer(b)
	var sign = md5Sign("POST", UMENG_MSG_SEND_API_URL, string(b), message.AppSecret)

	fmt.Println(string(b))

	req, err := http.NewRequest("POST", UMENG_MSG_SEND_API_URL+"?sign="+sign, buf)
	if err != nil {
		return nil
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")

	var client = http.DefaultClient
	rep, err := client.Do(req)
	if err != nil {
		return nil
	}
	defer rep.Body.Close()

	data, err := ioutil.ReadAll(rep.Body)
	if err != nil {
		return nil
	}
	json.Unmarshal(data, &results)

	return results
}

func md5Sign(method, url, postBody, appSecret string) string {
	var sign = fmt.Sprintf("%s%s%s%s", method, url, postBody, appSecret)
	var m = md5.New()
	m.Write([]byte(sign))
	sign = hex.EncodeToString(m.Sum(nil))
	return sign
}