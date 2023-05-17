package errs

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
	"testing"
)

func TestHandleError(t *testing.T) {
	type args struct {
		err    error
		log    *logrus.Logger
		tgbot  *tgbotapi.BotAPI
		chatID int64
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "gg",
			args: args{
				err:    nil,
				log:    nil,
				tgbot:  nil,
				chatID: 2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			HandleError(tt.args.err, tt.args.log, tt.args.tgbot, tt.args.chatID)
		})
	}
}
