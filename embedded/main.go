package main

import "fmt"

type Pony interface {
	Walk()
	Sprint()
}

type EarthPony struct {
}

func (ep *EarthPony) Walk() {
	fmt.Println("Walk")
}

func (ep *EarthPony) Sprint() {
	ep.Walk()
	ep.Walk()
	ep.Walk()
}

type Pegasus struct {
	EarthPony
}

func (u *Pegasus) Fly() {
	fmt.Println("Run")
}

func (u *Pegasus) Walk() {
	u.Fly()
}

func main() {
	ep := EarthPony{}
	ep.Walk()
	u := Pegasus{}
	u.Walk()
}
