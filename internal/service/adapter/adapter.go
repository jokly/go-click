package adapter

type Adapter interface {
	Send(event interface{}) error
}
