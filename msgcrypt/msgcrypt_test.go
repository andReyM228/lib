package msgcrypt

import (
	"testing"
)

func TestEncryptor_Encrypt_Decrypt(t *testing.T) {
	e := Encryptor{}
	key, err := e.GenerateKey(64)
	if err != nil {
		return
	}
	type fields struct {
		key []byte
	}
	type args struct {
		text string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				key: key,
			},
			args: args{
				text: "example",
			},
			want:    "example",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Encryptor{
				key: tt.fields.key,
			}
			eText, err := e.Encrypt(tt.args.text)
			if (err != nil) != tt.wantErr {
				t.Errorf("Encrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			got, err := e.Decrypt(eText)
			if (err != nil) != tt.wantErr {
				t.Errorf("Encrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got != tt.want {
				t.Errorf("Encrypt() got = %v, want %v", got, tt.want)
			}
		})
	}
}
