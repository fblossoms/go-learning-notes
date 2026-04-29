package main

import "fmt"

type Person3 struct {
}

type Runner interface {
	run()
}

func (Person3) run() {
}

func (Person3) swim() {
}

func mian() {
	t := Person3{}
	fmt.Println(t)
	t.run()
	var r Runner = t
	r.run()
	//r.swim()
}
