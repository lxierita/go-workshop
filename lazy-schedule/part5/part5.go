package part5

import "sync"

type intfunc func(...int) int

//Scheduler contains the scheduled functions
type Scheduler struct {
	scheduled []intfunc
}

//NewScheduler creates a new scheduler instance
func NewScheduler() *Scheduler {
	s := Scheduler{}
	return &s
}

//Add functions to scheduler lazily
func (s *Scheduler) Add(f intfunc, args []int) {
	af := func(...int) int {
		return f(args...)
	}
	s.scheduled = append(s.scheduled, af)
}

type jobs struct {
	index  int
	result int
}

//Run scheduled functions in the order that they were added
func (s *Scheduler) Run() []int {
	var wg sync.WaitGroup
	wg.Add(len(s.scheduled))

	r := make([]int, len(s.scheduled))
	ch := make(chan jobs, 1000000)
	done := make(chan bool)

	for i, f := range s.scheduled {
		go func(i int, f intfunc) {
			ch <- jobs{i, f()}
			if last := len(s.scheduled) - 1; i == last {
				wg.Wait()
				defer close(done)
				done <- true
			}
		}(i, f)
	}

	for {
		select {
		case j := <-ch:
			r[j.index] = j.result
			wg.Done()
		case <-done:
			close(ch)
			return r
		}
	}

}
