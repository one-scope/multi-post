package discord

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
)

type credentials struct {
	Token   string `json:"Token"`
	GuildID string `json:"GuildID"`
}

type Bot struct {
	*discordgo.Session
	Channels []*discordgo.Channel
}

func (aDiscord *Bot) SetCredentials(aFile string) error {
	tByte, tError := os.ReadFile(aFile)
	if tError != nil {
		return fmt.Errorf("failed to read %s: %w", aFile, tError)
	}
	tCredentials := credentials{}
	if tError := json.Unmarshal(tByte, &tCredentials); tError != nil {
		return fmt.Errorf("failed to extract credentials from %s: %w", aFile, tError)
	}
	tToken := "Bot " + tCredentials.Token
	if aDiscord.Session, tError = discordgo.New(tToken); tError != nil {
		return fmt.Errorf("failed to establish session: %w", tError)
	}
	if aDiscord.Channels, tError = aDiscord.Session.GuildChannels(tCredentials.GuildID); tError != nil {
		return fmt.Errorf("failed to load channels: %w", tError)
	}
	return nil
}

func (aDiscord Bot) SendMessage(aChannel string, aContent string) error {
	for _, channel := range aDiscord.Channels {
		if channel.Name != aChannel {
			continue
		}
		if _, tError := aDiscord.Session.ChannelMessageSend(channel.ID, aContent); tError != nil {
			return fmt.Errorf("failed to send message to discord: %w", tError)
		}
		return nil
	}
	return fmt.Errorf("failed to send message to discord: channel_not_found")
}
