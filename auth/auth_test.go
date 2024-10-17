package auth

import "testing"

func TestCreateToken(t *testing.T) {
	type args struct {
		chatID int64
		userID int64
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "succes create token",
			args: args{
				chatID: 1,
				userID: 5,
			},
			want:    "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODM5MDcyMjIsImlhdCI6MTY4MzkwNjYyMiwic3ViIjoxLCJ1c2VyX2lkIjo1fQ.IYHrvhJ3D9w_YGCl7T8uRmdnPgbLq59LY_ZwAOn8gA4",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateToken(tt.args.chatID, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CreateToken() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVerifyToken(t *testing.T) {
	type args struct {
		tokenString string
		chatID      []int64
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "succes verify token",
			args: args{
				tokenString: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODM5MDcyMjIsImlhdCI6MTY4MzkwNjYyMiwic3ViIjoxLCJ1c2VyX2lkIjo1fQ.IYHrvhJ3D9w_YGCl7T8uRmdnPgbLq59LY_ZwAOn8gA4",
				chatID:      []int64{5},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := VerifyToken(tt.args.tokenString, tt.args.chatID...); (err != nil) != tt.wantErr {
				t.Errorf("VerifyToken() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
