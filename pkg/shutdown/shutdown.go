package shutdown

import (
	"os"
	"os/signal"
	"syscall"
)

type Hook interface {
	OnShutdown(funcs ...func())
}

type hook struct {
	ch chan os.Signal
}

func New(signals ...syscall.Signal) Hook {
	h := &hook{
		ch: make(chan os.Signal, 1),
	}

	for _, s := range signals {
		signal.Notify(h.ch, s)
	}
	return h
}

func (h *hook) OnShutdown(funcs ...func()) {
	<-h.ch

	for _, f := range funcs {
		f()
	}

	close(h.ch)
}
