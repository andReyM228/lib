package discord

import (
	"context"
	"errors"
	"github.com/bwmarrin/discordgo"
	"log"
)

type discord struct {
	config       Config
	dsConnection *discordgo.Session
}

func Init(ctx context.Context, cfg Config) Discord {
	dsConn, err := discordgo.New("Bot " + cfg.Token)
	if err != nil {
		log.Fatal(err)
	}

	ds := discord{
		config:       cfg,
		dsConnection: dsConn,
	}

	ds.serveDiscord(ctx)

	return ds
}

func (d discord) Send(opts ...Option) error {
	if opts == nil {
		return errors.New("options must be")
	}

	msg := &discordgo.MessageSend{}

	for _, opt := range opts {
		opt(msg)
	}

	_, err := d.dsConnection.ChannelMessageSendComplex(d.config.TextChanelId, msg)
	if err != nil {
		return err
	}

	return nil
}

// serveDiscord open discord session with graceful shutdown.
func (d discord) serveDiscord(ctx context.Context) {
	if err := d.dsConnection.Open(); err != nil {
		log.Fatalf("🔴 discord: can`t open the session: %s", err)
	}

	log.Println("🟢 discord bot is ready")

	// graceful shutdown listener.
	go func() {
		<-ctx.Done()

		if err := d.dsConnection.Close(); err != nil {
			log.Fatalf("🔴 discord: can`t close the session: %s", err)
		}
	}()
}
