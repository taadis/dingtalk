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

func TestMain(m *testing.M) {
	testAccessToken = env.GetAccessToken()
	testSecret = env.GetSecret()
	os.Exit(m.Run())
}
