package main

import (
	"fmt"
	"io"
	"time"
)

type Pool struct {
	io.Reader
	P   chan []byte
	Max int
}

func NewPool(size int) *Pool {
	return &Pool{
		P:   make(chan []byte, size),
		Max: size,
	}
}

func (p *Pool) Pop() []byte {
	if len(p.P) == 0 {
		return nil
	}
	fmt.Println("leng pop p.p == ", len(p.P))
	return <-p.P
}

func (p *Pool) Put(b []byte) {

	fmt.Println("leng put p.p == ", len(p.P))
	if len(p.P) >= p.Max {
		return
	}
	p.P <- b
}

func Produce(ch chan<- []byte, j int) {
	for i := 0; i < j; i++ {
		ch <- []byte{byte(i)}
		fmt.Println("p=", i)
	}
	time.Sleep(time.Microsecond)
}

func Consumer(ch <-chan []byte) {
	for {
		v := <-ch
		fmt.Println("c=", v)
	}
	time.Sleep(time.Microsecond)
}

func Pt() {
	p := NewPool(1024)

	go Produce(p.P, 2)
	//go Consumer(p.P)

	for {
		v := <-p.P
		fmt.Println("c=", v)
		//if len(p.P) == 0 {
		//	break
		//}
	}
	defer close(p.P)
	//	time.Sleep(time.Second)

}
