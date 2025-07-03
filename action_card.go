package dingtalk

import (
	"context"
	"encoding/json"
)

type ActionCardMessage struct {
	//
	MsgType MsgType `json:"msgtype,omitempty"`
	//
	ActionCard ActionCardModel `json:"actionCard,omitempty"`
}

type ActionCardModel struct {
	//
	Title string `json:"title"`
	//
	Text string `json:"text"`
	// 单个按钮的标题
	SingleTitle string `json:"singleTitle"`
	// 点击消息跳转的URL
	SingleUrl string `json:"singleURL"`
	// 按钮排列方式
	// - 0：按钮竖直排列
	// - 1：按钮横向排列
	BtnOrientation string `json:"btnOrientation,omitempty"`
}

func (m *ActionCardMessage) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

func NewActionCardMessage(title string, text string, singleTitle string, singleUrl string) *ActionCardMessage {
	msg := new(ActionCardMessage)
	msg.MsgType = MsgTypeActionCard
	msg.ActionCard.Title = title
	msg.ActionCard.Text = text
	msg.ActionCard.SingleTitle = singleTitle
	msg.ActionCard.SingleUrl = singleUrl

	return msg
}

func (c *Client) SendActionCardMessage(ctx context.Context, msg *ActionCardMessage) error {
	return c.send(ctx, msg)
}
