package models

type Sender interface {
	Send(payload *Payload) error
}
