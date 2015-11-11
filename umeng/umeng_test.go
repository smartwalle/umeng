package umeng

import (
	"testing"
	"fmt"
)

func Test_iOS(t *testing.T) {
	var payload = NewUMengiOSPayload()
	payload.SetAlert("这是我发出来的消息哦board")
	payload.SetBadge(11)
	payload.Set("url", "http://www.baidu.com")
	payload.Set("type", "8")
	fmt.Println(PushBroadcastMessage("app key", "app secret", false, payload))
}

func Test_Android(t *testing.T) {
	var payload = NewUMengAndroidPayload()
	payload.SetDisplayType(UMENG_ANDROID_DISPLAY_TYPE_OF_NOTIFICATION)
	payload.SetTicker("ticker")
	payload.SetTitle("message")
	payload.SetText("text")
	payload.SetPlayLights(true)
	payload.SetPlaySound(true)
	payload.SetPlayVibrate(true)
	payload.SetAfterOpen("go_url", "http://www.baidu.com")

	var message = NewUMengMessage("app key", "app secret", UMENG_MESSAGE_TYPE_UNICAST, false)
	message.DeviceTokens = "device token"
	message.Payload = payload

	fmt.Println(PushMessage(message))
}