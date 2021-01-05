package main

import (
	"github.com/jormin/go-study/modules/log"
)

type A struct {
	B *int
	C *int
}

func main() {
	var b int = 0
	a := &A{
		B: &b,
	}
	log.Info("a: %+v", *a)
	var c int = 1
	a.C = &c
	log.Info("a: %+v", *a)
}
