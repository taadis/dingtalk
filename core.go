package dingtalk

import "context"

const Webhook = "https://oapi.dingtalk.com/robot/send?"

type Messager interface {
	Marshal() ([]byte, error)
}

type Sender interface {
	Send(ctx context.Context, msg Messager) error
}

type MsgType string

const (
	MsgTypeText       MsgType = "text"
	MsgTypeLink       MsgType = "link"
	MsgTypeMarkdown   MsgType = "markdown"
	MsgTypeActionCard MsgType = "actionCard"
	MsgTypeFeedCard   MsgType = "feedCard"
)

type AtModel struct {
	//
	AtMobiles []string `json:"atMobiles,omitempty"`
	//
	AtUserIds []string `json:"atUserIds,omitempty"`
	//
	IsAtAll bool `json:"isAtAll,omitempty"`
}

type Response struct {
	// error code
	ErrCode int `json:"errcode"`
	// error message
	ErrMsg string `json:"errmsg"`
}
