package dingtalk

import (
	"context"
	"testing"
)

func TestSendTextMessage(t *testing.T) {
	ctx := context.Background()

	msg := NewTextMessage("test text message...", WithAtAll(true), WithAtMobiles([]string{}))

	accessToken := WithAccessToken(testAccessToken)
	secret := WithSecret(testSecret)
	client := NewClient(accessToken, secret)
	err := client.SendTextMessage(ctx, msg)
	if err != nil {
		t.Fatal(err)
	}
}
