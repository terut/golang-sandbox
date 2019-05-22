package main

import (
	"fmt"
	"sync"
	"time"
)

type sample struct {
	str string
}

type State struct {
	sm sync.Map
}

func (s *State) Add(key string, value *sample) {
	s.sm.Store(key, value)
}

func (s *State) Get(key string) *sample {
	v, _ := s.sm.Load(key)
	sp := v.(*sample)
	return sp
}

func main() {
	s := &sample{str: "hoge"}
	state := new(State)
	state.Add("hoge", s)

	go func() {
		fmt.Println("start func1")

		ns := state.Get("hoge")
		<-time.After(10 * time.Second)
		//ns.str = "fuga"
		fmt.Printf("func1 str: %s\n", ns.str)

		fmt.Printf("func1 pointer: %p\n", ns)
		fmt.Println("end func1")
	}()

	go func() {
		fmt.Println("start func2")

		ns := state.Get("hoge")
		<-time.After(5 * time.Second)
		//ns.str = "fuga"
		fmt.Printf("func2 str: %s\n", ns.str)

		fmt.Printf("func2 pointer: %p\n", ns)
		fmt.Println("end func2")
	}()

	for {
	}
}
