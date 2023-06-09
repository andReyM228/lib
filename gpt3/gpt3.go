package gpt3

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func getCompletion(prompt string, apiKey string) (string, error) {
	url := "https://api.openai.com/v1/engines/davinci-codex/completions"

	data := map[string]interface{}{
		"prompt":     prompt,
		"max_tokens": 10,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", fmt.Errorf("Error encoding JSON: %v", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("Error creating request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("Error sending request: %v", err)
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("Error reading response body: %v", err)
	}

	var result map[string]interface{}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return "", fmt.Errorf("Error decoding response JSON: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Request failed with status code %d: %s", resp.StatusCode, responseBody)
	}

	if output, ok := result["choices"].([]interface{}); ok && len(output) > 0 {
		if choice, ok := output[0].(map[string]interface{}); ok {
			if text, ok := choice["text"].(string); ok {
				return text, nil
			}
		}
	}

	return "", fmt.Errorf("Unexpected response format: %s", responseBody)
}

func main() {
	prompt := "Hello,"
	apiKey := "YOUR_API_KEY"

	response, err := getCompletion(prompt, apiKey)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(response)
}
