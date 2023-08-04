package sheets

import (
	"context"
	"embed"
)

const (
	credentialsFile = "credentials.json"
)

func initSheets(ctx context.Context, cred embed.FS) error {
	credFile, err := cred.ReadFile(credentialsFile)
	if err != nil {
		return err
	}

	return nil
}
