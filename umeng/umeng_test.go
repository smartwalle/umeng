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
	fmt.Println(PushBroadcastMessage("55a37b3c67e58ecf6d001bc5", "rvs9vjiqrz0eyzqtdc9qz2qswriglgpz", false, payload))
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

	var message = NewUMengMessage("53ec8886fd98c587cc001ff5", "991b07edd5d1d6a8a8d1d053f0763b70", UMENG_MESSAGE_TYPE_UNICAST, false)
	message.DeviceTokens = "Ag-YlpXmICyhAhsEMrNMJhV-KQjSQCY61D8j1izNxkj3"
	message.Payload = payload

	fmt.Println(PushMessage(message))
}