package shutdown

import (
	"os"
	"os/signal"
	"sync"
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

	wg := &sync.WaitGroup{}
	for _, f := range funcs {
		wg.Add(1)
		go func(f func()) {
			defer wg.Done()
			f()
		}(f)
	}

	wg.Wait()
	close(h.ch)
}
