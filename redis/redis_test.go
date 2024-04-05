package redis

import (
	"encoding/json"
	"github.com/andReyM228/lib/log"
	"testing"
)

type user struct {
	Name    string
	Surname string
	Age     int
}

func Test_cache_Set(t *testing.T) {
	c := NewCache(Config{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	},
		log.Init(),
	)

	userBytes, err := json.Marshal(user{
		Name:    "test",
		Surname: "test",
		Age:     10,
	})
	if err != nil {
		t.Error(err)
	}

	type args struct {
		key   string
		value interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "success",
			args: args{
				key:   "user1",
				value: userBytes,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := c.Set(tt.args.key, tt.args.value); err != nil {
				t.Errorf("Set() error = %v", err)
			}
		})
	}
}

func Test_cache_GetBytes(t *testing.T) {
	c := NewCache(Config{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	},
		log.Init(),
	)

	type args struct {
		key string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "success",
			args: args{
				key: "user1",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := c.GetBytes(tt.args.key)
			if err != nil {
				t.Errorf("GetBytes() error = %v", err)
				return
			}

			var user user
			err = json.Unmarshal(got, &user)
			if err != nil {
				t.Error(err)
			}

			t.Logf("%+v", user)
		})
	}
}
