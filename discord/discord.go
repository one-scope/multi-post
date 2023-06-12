package discord

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

var (
	Token    = "MTExNDg2MDI3NzYwMzc3ODYwMA.GIlUuP.yecJsyGp16z-EHsuCs_eiFKI3oP73WPeNbClOI" //"Bot"という接頭辞がないと401 unauthorizedエラーが起きます
	ClientID = "1114860277603778600"
	GuildID  = "1114864523350900769" //サーバID
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
