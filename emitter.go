package emitter

type Emitter struct {
	eMap map[string][]Listener
}

type Listener struct {
	key string
	f   *func()
}

// Subscribe on event
func (e *Emitter) On(event string, key string, cb func()) {
	if _, ok := e.eMap[event]; !ok {
		e.eMap[event] = make([]Listener, 0)
	}

	e.eMap[event] = append(e.eMap[event], Listener{key: key, f: &cb})
}

// Emit event by name and call all listeners
func (e *Emitter) Emit(event string) {
	if listeners, ok := e.eMap[event]; ok {
		if listeners != nil {
			for _, listener := range listeners {
				(*listener.f)()
			}
		}
	}
}

func (e *Emitter) RemoveListener(event string, key string) {
	if listeners, ok := e.eMap[event]; ok {
		for i, listener := range listeners {
			if listener.key == key {
				e.eMap[event] = append(e.eMap[event][:i], e.eMap[event][i+1:]...)
			}
		}
	}
}

func (e *Emitter) RemoveAllListeners(event string) {
	delete(e.eMap, event)
}

func Create() *Emitter {
	return &Emitter{make(map[string][]Listener)}
}
