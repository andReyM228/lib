package discord

type Discord interface {
	Send(opts ...Option) error
}
