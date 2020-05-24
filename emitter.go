package emitter

type Emitter struct {
	eMap map[string][]func()
}

func (e *Emitter) On(event string, cb func()) {
	if _, ok := e.eMap[event]; !ok {
		e.eMap[event] = make([]func(), 0)
	}

	e.eMap[event] = append(e.eMap[event], cb)
}

func (e *Emitter) Emit(event string) {
	var listeners *[]func()

	for key, cbs := range e.eMap {
		if key == event {
			listeners = &cbs
		}
	}

	if listeners != nil {
		for _, cb := range *listeners {
			cb()
		}
	}
}

func Create() *Emitter {
	return &Emitter{make(map[string][]func())}
}
