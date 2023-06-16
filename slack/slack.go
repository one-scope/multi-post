package slack

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/slack-go/slack"
)

type credentials struct {
	Token string `json:"Token"`
}

type Bot struct {
	*slack.Client
}

func (aSlack *Bot) SetCredentials(aFile string) error {
	tByte, tError := os.ReadFile(aFile)
	if tError != nil {
		return fmt.Errorf("failed to read %s: %w", aFile, tError)
	}
	tCredentials := credentials{}
	if tError := json.Unmarshal(tByte, &tCredentials); tError != nil {
		return fmt.Errorf("failed to extract credentials from %s: %w", aFile, tError)
	}
	aSlack.Client = slack.New(tCredentials.Token)
	return nil
}

func (aSlack Bot) SendMessage(aChannel string, aContent string) error {
	if _, _, tError := aSlack.Client.PostMessage(aChannel, slack.MsgOptionText(aContent, true)); tError != nil {
		return fmt.Errorf("failed to send message to slack: %w", tError)
	}
	return nil
}
