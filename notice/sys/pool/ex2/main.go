package main

import (
	"fmt"
	"sync"
)

func main() {
	p := &sync.Pool{}
	a := p.Get()
	if a == nil {
		a = func() interface{} {
			return 0
		}
	}
	p.Put(1)
	b := p.Get().(int)
	fmt.Println(a, b)
}
