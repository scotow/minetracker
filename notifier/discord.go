package notifier

import (
	"github.com/bwmarrin/discordgo"
)

func NewDiscordNotifier(token, channel string) (*DiscordNotifier, error) {
	session, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, err
	}

	err = session.Open()
	if err != nil {
		return nil, err
	}

	session.ShouldReconnectOnError = true

	dn := new(DiscordNotifier)
	dn.session = session
	dn.channel = channel

	return dn, nil
}

type DiscordNotifier struct {
	session *discordgo.Session
	channel string
}

func (dn *DiscordNotifier) Notify(message string) error {
	_, err := dn.session.ChannelMessage(dn.channel, message)
	return err
}
