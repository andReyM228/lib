package gpt3

import "testing"

func Test_gpt3_GetCompletion(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "",
			args: args{
				text: "2+2?",
			},
			want:    "",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := Init("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", "gpt-3.5-turbo")
			got, err := g.GetCompletion(tt.args.text)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCompletion() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetCompletion() got = %v, want %v", got, tt.want)
			}
		})
	}
}
