package database

import (
	"embed"
	"github.com/andReyM228/lib/log"
	"github.com/jmoiron/sqlx"
	"testing"
)

//go:embed fixtures/migrations
var dbMigrationFS embed.FS

func TestInitDatabase(t *testing.T) {

	type args struct {
		log    log.Logger
		config DBConfig
	}
	tests := []struct {
		name string
		args args
		want *sqlx.DB
	}{
		{
			name: "success",
			args: args{
				log: log.Init(),
				config: DBConfig{
					Host:            "localhost",
					Port:            5432,
					User:            "postgres",
					Password:        "postgres",
					DBName:          "test",
					ConfigDirectory: "fixtures/migrations",
					AutoMigrate:     true,
				},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := InitDatabase(tt.args.log, tt.args.config, dbMigrationFS)
			t.Log(got)
		})
	}
}
