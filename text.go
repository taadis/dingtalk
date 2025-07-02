package dingtalk

import (
	"context"
	"encoding/json"
)

type TextModel struct {
	Content string `json:"content,omitempty"`
}

type TextMessage struct {
	//
	MsgType MsgType `json:"msgtype,omitempty"`
	//
	Text TextModel `json:"text,omitempty"`
	//
	At AtModel `json:"at,omitempty"`
}

func (m *TextMessage) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

func NewTextMessage(content string, opts ...AtOption) *TextMessage {
	msg := new(TextMessage)
	msg.MsgType = MsgTypeText
	msg.Text.Content = content

	for _, opt := range opts {
		opt.Apply(&msg.At)
	}

	return msg
}

func (c *Client) SendTextMessage(ctx context.Context, msg *TextMessage) error {
	return c.send(ctx, msg)
}
