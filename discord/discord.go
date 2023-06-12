package discord

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

var (
	Token   = "xxxxxxxx" //"Bot"という接頭辞がないと401 unauthorizedエラーが起きます
	GuildID = "xxxxxxxx" //サーバID
)

type Bot struct {
	*discordgo.Session
	Channels []*discordgo.Channel
}

func (aDiscord *Bot) SetCredentials(aFIle string) {
	aDiscord.Session, _ = discordgo.New("Bot" + Token)
	aDiscord.Channels, _ = aDiscord.Session.GuildChannels(GuildID)

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
	return fmt.Errorf("the specified channel cannot be found")
}
