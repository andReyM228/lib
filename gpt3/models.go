package gpt3

type chatResponse struct {
	ID      string   `json:"id"`
	Choices []choice `json:"choices"`
}

type choice struct {
	Message message `json:"message"`
}

type message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}
