package dingtalk

import (
	"context"
	"encoding/json"
)

type LinkMessage struct {
	//
	MsgType MsgType `json:"msgtype,omitempty"`
	//
	Link LinkModel `json:"link,omitempty"`
}

type LinkModel struct {
	// title
	Title string `json:"title"`
	// text
	Text string `json:"text"`
	//
	MessageUrl string `json:"messageUrl"`
	//
	PicUrl string `json:"picUrl,omitempty"`
}

func (m *LinkMessage) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

func NewLinkMessage(title string, text string, messageUrl string) *LinkMessage {
	msg := new(LinkMessage)
	msg.MsgType = MsgTypeLink
	msg.Link.Title = title
	msg.Link.Text = text
	msg.Link.MessageUrl = messageUrl

	return msg
}

func (c *Client) SendLinkMessage(ctx context.Context, msg *LinkMessage) error {
	return c.send(ctx, msg)
}
