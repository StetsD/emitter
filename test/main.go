package main

import (
	"emitter"
	"fmt"
)

type Worker struct {
	name string
}

func (w *Worker) Work() {
	fmt.Println("I am " + w.name + " and I work msg: ")
}

func (w *Worker) Notify() {
	fmt.Println(w.name, " says TurboLove")
}

type Node struct {
	name string
	emitter.Emitter
}

func (n *Node) message(msg string) {
	fmt.Println("I am "+n.name+" and I receive msg: ", msg)
}

func main() {
	event1 := "superKillMazafakaSignal"
	worker1 := Worker{"worker_1"}
	worker2 := Worker{"worker_2"}
	node1 := Node{"node_1", *emitter.Create()}

	node1.On(event1, "worker1.Work", worker1.Work)
	node1.On(event1, "worker2.Work", worker2.Work)
	node1.On(event1, "worker1.Notify", worker1.Notify)
	node1.RemoveListener(event1, "worker2.Work")
	//

	node1.Emit(event1)
	node1.RemoveListener(event1, "worker1.Notify")
	node1.Emit(event1)
	node1.RemoveAllListeners(event1)
	node1.Emit(event1)
}
