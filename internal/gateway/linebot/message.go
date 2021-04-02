package linebot

import "github.com/line/line-bot-sdk-go/linebot"

type callMessagingApi struct {
	client *linebot.Client
}

func newMessagingApi(channelSecret, channelToken string) (*callMessagingApi, error) {
	bot, err := linebot.New(channelSecret, channelToken)
	if err != nil {
		return nil, err
	}
	call := &callMessagingApi{
		client: bot,
	}
	return call, nil
}

type callPushMessage struct {
	client   *linebot.Client
	to       string
	messages []linebot.SendingMessage
}

func (c *callMessagingApi) newPushMessage(to string, messages ...linebot.SendingMessage) *callPushMessage {
	return &callPushMessage{
		client:   c.client,
		to:       to,
		messages: messages,
	}
}

func (c *callPushMessage) do() error {
	if _, err := c.client.PushMessage(c.to, c.messages...).Do(); err != nil {
		return err
	}
	return nil
}

type callFlexMessage struct {
	altText string
	data    []byte
}

func newFlexMessage(altText string, data []byte) *callFlexMessage {
	return &callFlexMessage{
		altText: altText,
		data:    data,
	}
}

func (c *callFlexMessage) do() (*linebot.FlexMessage, error) {
	contents, err := linebot.UnmarshalFlexMessageJSON(c.data)
	if err != nil {
		return nil, err
	}
	return linebot.NewFlexMessage(c.altText, contents), nil
}
