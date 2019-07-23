package util

type Semaphore struct {
	c chan bool
}

func NewSemaphore(n int) *Semaphore {
	return &Semaphore{make(chan bool, n)}
}

func (s *Semaphore) Acquire() {
	s.c <- true
}

func (s *Semaphore) Release() {
	<-s.c
}
