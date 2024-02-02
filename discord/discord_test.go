package discord

import (
	"context"
	"testing"
)

func Test_discord_Send(t *testing.T) {

	//file, err := os.Open("test-video.mp4")
	//if err != nil {
	//	t.Log("error opening file: ", err)
	//	return
	//}

	d := Init(context.Background(), Config{
		Token:        "MTIwMDQ5NDM2NDEyMDU4NDQwMg.Ghw_mv.ernVQ2dDX-nglzibamysi_-fTfZdIQ08n9jUOc",
		TextChanelId: "1014087455101694004",
		BotId:        "1200494364120584402",
	})

	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "success",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := d.Send(WithMessage("test test"), WithButton("button", "https://meet.google.com/adi-hjqr-bew"), WithButton("button", "https://meet.google.com/adi-hjqr-bew")); err != nil {
				t.Errorf("Send() error = %v", err)
			}
		})
	}
}
