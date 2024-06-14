package eventmachine

type EventPipeIfaces interface {
	Push(event string, args ...interface{})
	Pull() (event string, args []interface{})
}

type WhenIface interface {
	When(event string, handler func(...interface{}))
}
type EventMachine interface {
	WhenIface
}

type EventPipe struct {
}

func NewEventPipe() *EventPipe {
	return &EventPipe{}
}

func (ep *EventPipe) Push(event string, args ...interface{}) {

}


func (ep *EventPipe) Pull() (event string, args []interface{}) {
	return
}


