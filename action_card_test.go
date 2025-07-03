package dingtalk

import (
	"context"
	"testing"
)

func TestSendActionCardMessage(t *testing.T) {
	ctx := context.Background()

	title := "乔布斯 20 年前想打造一间苹果咖啡厅，而它正是 Apple Store 的前身"
	text := "链接类型消息"
	singleTitle := "ActionCard消息类型"
	singleUrl := testUrl
	msg := NewActionCardMessage(title, text, singleTitle, singleUrl)
	msg.ActionCard.BtnOrientation = "0"
	accessToken := WithAccessToken(testAccessToken)
	secret := WithSecret(testSecret)
	client := NewClient(accessToken, secret)
	err := client.SendActionCardMessage(ctx, msg)
	if err != nil {
		t.Fatal(err)
	}
}
