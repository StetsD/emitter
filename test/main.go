package main

import (
	"emitter"
	"fmt"
)

type Worker struct {
	name string
}

func (w *Worker) work() {
	fmt.Println("I am " + w.name + " and I work msg: ")
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

	node1.On(event1, worker1.work)
	node1.On(event1, worker2.work)

	node1.Emit(event1)
}
