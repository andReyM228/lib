package sheets

import (
	"context"
	"embed"
	"testing"
)

//go:embed fixtures
var credFile embed.FS

func Test_googleSheets_Get(t *testing.T) {
	g, err := initSheets(context.Background(), credFile)
	if err != nil {
		t.Fatal(err)
	}
	tests := []struct {
		name string
	}{
		{
			name: "success",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g = g.WithSheetID("1YTOSQP4TGYjgrpRdaGBrsKbxGs1jCzFBi0eB-TZzDUM").WithReadRange("Лист1!A1:B3")
			g.Get()
		})
	}
}
