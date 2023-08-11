package rabbit

import (
	"github.com/streadway/amqp"
	"testing"
)

func TestRabbitMQ_NewRabbitMQ(t *testing.T) {
	type fields struct {
		conn *amqp.Connection
		ch   *amqp.Channel
	}
	type args struct {
		url string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "",
			fields: fields{},
			args: args{
				url: "amqp://guest:guest@localhost:5672/",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &RabbitMQ{
				conn: tt.fields.conn,
				ch:   tt.fields.ch,
			}
			if err := r.NewRabbitMQ(tt.args.url); (err != nil) != tt.wantErr {
				t.Errorf("NewRabbitMQ() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
