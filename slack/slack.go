package slack

import (
	"fmt"

	"github.com/slack-go/slack"
)

var token = "adsfadfadf"

type Bot struct {
	*slack.Client
}

func (aSlack *Bot) SetCredentials(aFile string) {
	aSlack.Client = slack.New(token)
}

func (aSlack Bot) SendMessage(aChannel string, aContent string) error {
	if _, _, tError := aSlack.Client.PostMessage(aChannel, slack.MsgOptionText(aContent, true)); tError != nil {
		return fmt.Errorf("failed to send message to slack: %w", tError)
	}
	return nil
}
