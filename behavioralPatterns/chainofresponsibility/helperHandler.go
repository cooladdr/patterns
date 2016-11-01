package chainofresponsibility

import (
	_ "fmt"
)

type Topic int

const (
	NO_HELP_TOPIC = Topic(-1)
)

type HelpHandler struct {
	successor *HelpHandler
	topic     Topic
}

func (h *HelpHandler) HasHelp() bool {
	return h.topic != NO_HELP_TOPIC
}

func (h *HelpHandler) HandleHelp() {
	if h.successor != nil {
		h.successor.HandleHelp()
	}
}

func (h *HelpHandler) SetHandler(hp *HelpHandler, t Topic) {
	h.successor = hp
	h.topic = t
}

type Widget struct {
	HelpHandler
	parent *Widget
}

type Button struct {
	Widget
}

func (b *Button) HandleHelp() {
	b.HelpHandler.HandleHelp()
}
