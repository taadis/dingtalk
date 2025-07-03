package dingtalk

import (
	"os"
	"testing"

	"github.com/taadis/dingtalk/env"
)

// test access token
var testAccessToken string

// test secret
var testSecret string

// test url
var testUrl string

func TestMain(m *testing.M) {
	testAccessToken = env.GetAccessToken()
	testSecret = env.GetSecret()
	testUrl = "https://open.dingtalk.com/document/orgapp/custom-bot-send-message-type"
	os.Exit(m.Run())
}
