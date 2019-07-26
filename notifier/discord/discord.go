package discord

import (
	"github.com/bwmarrin/discordgo"
)

// Create a Discord Notifier.
// token is the connection token provided by the Discord Developers Portal. channel is the channel where the messages will be published.
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

// A Discord notifier. This Notifier send the message on the specified server channel.
type DiscordNotifier struct {
	session *discordgo.Session
	channel string
}

func (dn *DiscordNotifier) Notify(message string) error {
	_, err := dn.session.ChannelMessageSend(dn.channel, message)
	return err
}
