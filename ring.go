package gencode

import (
	"time"
)

func NewRing(min, max int, wait time.Duration) *Ring {
	r := &Ring{
		ch:   make(chan int),
		Min:  min,
		Max:  max,
		Wait: wait,
	}
	r.init()
	return r
}

type Ring struct {
	ch   chan int
	Min  int
	Max  int
	Wait time.Duration
}

func (r *Ring) init() {
	m := make(map[int]bool)
	for i := r.Min; i <= r.Max; i++ {
		m[i] = false
	}

	go func() {
		for {
			for k, _ := range m {
				r.Push(k)
			}
			//保证同一个周期内不会重复
			<-time.After(r.Wait)
		}
	}()
}

func (r *Ring) Next() int {
	return r.Pull()
}

func (r *Ring) Pull() int {
	return <-r.ch
}

func (r *Ring) Push(n int) {
	r.ch <- n
}