package services

type Params map[string]string

type INotification interface {
	Init(p *Params) (bool, error)
	Send(v any) error
}
