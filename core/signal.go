package core

type Signal interface {
	Signal()
	Wait()
}

type signalimpl struct {
	isOpen chan bool
}

func CreateSignal() Signal {
	return signalimpl{make(chan bool)}
}

func (s signalimpl) Signal() {
	close(s.isOpen)
}

func (s signalimpl) Wait() {
	for {
		select {
		case _, ok := <-s.isOpen:
			if !ok {
				return
			}
		}
	}
}
