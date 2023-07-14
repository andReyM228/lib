package gpt3

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	_ "github.com/andReyM228/lib/errs"
)

type gpt3 struct {
	apiKey string
	model  string
}

func Init(apiKey, model string) ChatGPT {
	return gpt3{
		apiKey: apiKey,
		model:  model,
	}
}

func (g gpt3) GetCompletion(text string) (string, error) {
	url := "https://api.openai.com/v1/chat/completions"

	data := fmt.Sprintf(`{
    	"model": "%s",
    	"messages": [{"role": "user", "content": "%s"}]
	}`, g.model, text)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(data)))
	if err != nil {
		return "", fmt.Errorf("Error creating request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", g.apiKey))

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("Error sending request: %v", err)
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("Error reading response body: %v", err)
	}

	var result chatResponse
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return "", fmt.Errorf("Error decoding response JSON: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Request failed with status code %d: %s", resp.StatusCode, responseBody)
	}

	if len(result.Choices) == 0 {
		return "", fmt.Errorf("Error")
	}

	return result.Choices[0].Message.Content, nil
}
