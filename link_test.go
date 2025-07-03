package dingtalk

import (
	"context"
	"testing"
)

func TestSendLinkMessage(t *testing.T) {
	ctx := context.Background()

	title := "自定义机器人发送消息的消息类型"
	text := "链接类型消息"
	messageUrl := "https://open.dingtalk.com/document/orgapp/custom-bot-send-message-type"
	msg := NewLinkMessage(title, text, messageUrl)
	msg.Link.PicUrl = ""
	accessToken := WithAccessToken(testAccessToken)
	secret := WithSecret(testSecret)
	client := NewClient(accessToken, secret)
	err := client.SendLinkMessage(ctx, msg)
	if err != nil {
		t.Fatal(err)
	}
}
