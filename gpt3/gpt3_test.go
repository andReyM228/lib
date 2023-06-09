package gpt3

import "testing"

func Test_getCompletion(t *testing.T) {
	type args struct {
		prompt string
		apiKey string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "test_1",
			args: args{
				prompt: "Hello,",
				apiKey: "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",
			},
			want:    "",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getCompletion(tt.args.prompt, tt.args.apiKey)
			if (err != nil) != tt.wantErr {
				t.Errorf("getCompletion() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getCompletion() got = %v, want %v", got, tt.want)
			}
		})
	}
}
