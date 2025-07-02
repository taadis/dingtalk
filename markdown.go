package dingtalk

import (
	"context"
	"encoding/json"
)

type MarkdownMessage struct {
	//
	MsgType MsgType `json:"msgtype,omitempty"`
	//
	Markdown MarkdownModel `json:"markdown,omitempty"`
	//
	At AtModel `json:"at,omitempty"`
}

type MarkdownModel struct {
	// title
	Title string `json:"title,omitempty"`
	// text
	Text string `json:"text,omitempty"`
}

func (m *MarkdownMessage) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

func NewMarkdownMessage(title string, text string, opts ...AtOption) *MarkdownMessage {
	msg := new(MarkdownMessage)
	msg.MsgType = MsgTypeMarkdown
	msg.Markdown.Title = title
	msg.Markdown.Text = text

	for _, opt := range opts {
		opt.Apply(&msg.At)
	}

	return msg
}

func (c *Client) SendMarkdownMessage(ctx context.Context, msg *MarkdownMessage) error {
	return c.send(ctx, msg)
}
