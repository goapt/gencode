package gencode

func NewRing(min, max int) *Ring {
	r := &Ring{
		ch:  make(chan int, max-min+1),
		Min: min,
		Max: max,
	}
	r.init()
	return r
}

type Ring struct {
	ch  chan int
	Min int
	Max int
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
	n := <-r.ch
	defer r.Push(n)
	return n
}

func (r *Ring) Pull() int {
	return <-r.ch
}

func (r *Ring) Push(n int) {
	r.ch <- n
}
