package part3

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

//Run scheduled functions in the order that they were added
func (s *Scheduler) Run() []int {
	var wg sync.WaitGroup

	r := make([]int, len(s.scheduled))
	wg.Add(len(s.scheduled))

	for i, f := range s.scheduled {
		go func(i int, f intfunc) {
			defer wg.Done()
			r[i] = f()
		}(i, f)
	}
	wg.Wait()
	return r
}
s