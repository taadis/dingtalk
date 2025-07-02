package dingtalk

import (
	"context"
	"testing"
)

func TestSendMarkdownMessage(t *testing.T) {
	ctx := context.Background()

	text := "# H1...\n ## H2...\n ### H3...\n markdown text..."
	msg := NewMarkdownMessage("test title", text, WithAtAll(true))
	accessToken := WithAccessToken(testAccessToken)
	secret := WithSecret(testSecret)
	client := NewClient(accessToken, secret)
	err := client.SendMarkdownMessage(ctx, msg)
	if err != nil {
		t.Fatal(err)
	}
}
