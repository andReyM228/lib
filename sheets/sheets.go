package sheets

import (
	"context"
	"embed"
	"encoding/json"
	"fmt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
	"log"
	"net/http"
	"os"
)

const (
	credentialsFile  = "fixtures/credentials1.json"
	defaultReadRange = "Sheet1!A1:Z100"
)

type googleSheets struct {
	client        *sheets.Service
	spreadsheetID string
	readRange     string
}

func getClient(config *oauth2.Config) *http.Client {
	// The file token.json stores the user's access and refresh tokens, and is
	// created automatically when the authorization flow completes for the first
	// time.
	tokFile := "token.json"
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(tokFile, tok)
	}
	return config.Client(context.Background(), tok)
}

func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code: %v", err)
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}
	return tok
}

func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

func initSheets(ctx context.Context, cred embed.FS) (GoogleSheets, error) {
	//credFile, err := cred.ReadFile(credentialsFile)
	//if err != nil {
	//	return googleSheets{}, err
	//}
	//
	//service, err := sheets.NewService(ctx, option.WithCredentialsJSON(credFile))
	//if err != nil {
	//	return googleSheets{}, err
	//}
	//
	//client := googleSheets{
	//	client:    service,
	//	readRange: defaultReadRange,
	//}
	//
	//return client, nil

	b, err := os.ReadFile("fixtures/credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, "https://www.googleapis.com/auth/spreadsheets.readonly")
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := getClient(config)

	srv, err := sheets.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}

	return googleSheets{
		client: srv,
	}, nil

}

func (g googleSheets) WithSheetID(sheetID string) GoogleSheets {
	g.spreadsheetID = sheetID

	return g
}

func (g googleSheets) WithReadRange(readRange string) GoogleSheets {
	g.readRange = readRange

	return g
}

func (g googleSheets) Get() {
	resp, err := g.client.Spreadsheets.Values.Get(g.spreadsheetID, g.readRange).Do()
	if err != nil {
		log.Fatalf("Una	ble to retrieve data from sheet: %v", err)
	}

	if len(resp.Values) == 0 {
		fmt.Println("No data found.")
		return
	}

	fmt.Println("Data:")
	for _, row := range resp.Values {
		fmt.Printf("%s, %s\n", row[0], row[1])
	}

}
