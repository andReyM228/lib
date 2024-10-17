package gpt3

type ChatGPT interface {
	GetCompletion(text string) (string, error)
}
