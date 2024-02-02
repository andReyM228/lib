package discord

import (
	"github.com/bwmarrin/discordgo"
	"io"
)

type Option func(msg *discordgo.MessageSend)

func WithMessage(text string) Option {
	return func(msg *discordgo.MessageSend) {
		msg.Content = text
	}
}

func WithFile(reader io.Reader, fileName string) Option {
	return func(msg *discordgo.MessageSend) {
		msg.Files = append(msg.Files, &discordgo.File{
			Name:   fileName,
			Reader: reader,
		})
	}
}

func WithButton(label, url string) Option {
	return func(msg *discordgo.MessageSend) {
		button := discordgo.Button{
			Label: label,
			Style: discordgo.LinkButton,
			Emoji: discordgo.ComponentEmoji{Name: "🔥"},
			URL:   url,
		}

		actionRow := discordgo.ActionsRow{
			Components: []discordgo.MessageComponent{button},
		}

		msg.Components = append(msg.Components, actionRow)
	}
}
