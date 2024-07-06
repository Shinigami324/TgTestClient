package events

type Featche interface {
	Featch(limit int) ([]Event, error)
}

type Processor interface {
	Process(e Event) error
}

type Type int

const (
	Unkown Type = iota
	Message
)

type Event struct {
	Type Type
	Text string
	Meta interface{}
}
