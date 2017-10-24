package common

type Message interface {
	Type() string
}
type MessageHandler func(Message)

type MessagingSystem struct {
	handlers map[string][]MessageHandler
}

func NewMessagingSystem() MessagingSystem {
	i := MessagingSystem{}
	i.handlers = map[string][]MessageHandler{}
	return i

}
func (ms *MessagingSystem) AddHandler(mesType string, mh MessageHandler) {
	if _, ok := ms.handlers[mesType]; !ok {
		ms.handlers[mesType] = []MessageHandler{}
	}
	ms.handlers[mesType] = append(ms.handlers[mesType], mh)
}
func (ms *MessagingSystem) Handle(m Message) {
	if _, ok := ms.handlers[m.Type()]; !ok {
		return
	}
	handlers := ms.handlers[m.Type()]
	for _, h := range handlers {
		h(m)
	}
}
