package gencode

import (
	"log"
	"time"
)

func NewRing(min, max int, wait time.Duration) *Ring {
	r := &Ring{
		ch:  make(chan int, max-min+1),
		Min: min,
		Max: max,
		Wait:wait,
	}
	return r
}

type Ring struct {
	ch  chan int
	Min int
	Max int
	Wait time.Duration
	preTime time.Time
}

func (r *Ring) init() {
	m := make(map[int]bool)
	for i := r.Min; i <= r.Max; i++ {
		m[i] = false
	}
	for k, _ := range m {
		r.ch <- k
	}
}

func (r *Ring) Next() int {
	n := time.Now()
	r.preTime = n

	for {
		select {
		case n := <-r.ch:
			return n
		default:
			log.Println("=====>init")
			//如果时间在同一秒，则需要暂停一秒，这样就可以保证不会在同一秒内生成重复的码
			if n.Sub(r.preTime) < r.Wait {
				<-time.After(r.Wait)
			}
			r.init()
		}
	}
}

func (r *Ring) Pull() int {
	return <-r.ch
}

func (r *Ring) Push(n int) {
	r.ch <- n
}
