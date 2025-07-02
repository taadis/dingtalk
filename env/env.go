package env

import "os"

func GetAccessToken() string {
	return os.Getenv("DINGTALK_ACCESS_TOKEN")
}

func GetSecret() string {
	return os.Getenv("DINGTALK_SECRET")
}
