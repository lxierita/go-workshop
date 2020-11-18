package part10

import (
	"log"
	"sync"
)

//Scheduler contains the scheduled functions
type Scheduler struct {
	scheduled []interface{}
}

//NewScheduler creates a new scheduler instance
func NewScheduler() *Scheduler {
	s := Scheduler{}
	return &s
}

//Add functions to scheduler lazily
func (s *Scheduler) Add(f interface{}) {
	af := func() interface{} {
		return f.(func() interface{})()
	}
	s.scheduled = append(s.scheduled, af)
}

type jobs struct {
	index  int
	result interface{}
}

//Job describes the execution result of each job
type Job struct {
	Result  interface{}
	Ok      bool
	Message string
}

//Run scheduled functions in the order that they were added
func (s *Scheduler) Run() (r []Job, err error) {

	var wg sync.WaitGroup
	wg.Add(len(s.scheduled))

	r = make([]Job, len(s.scheduled))
	ch := make(chan jobs, 10)
	done := make(chan bool)

	for i, f := range s.scheduled {
		defer func() {
			if err := recover(); err != nil {
				log.Println("in loop:  ", err)
				r[i] = Job{}
				wg.Done()
			}
		}()
		go func(i int, f interface{}) {
			defer func() {
				if err := recover(); err != nil {
					log.Println("in goroutines:  ", err)
					r[i] = Job{}
					wg.Done()
				}
			}()
			ch <- jobs{i, f.(func() interface{})()}
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
			r[j.index] = Job{j.result, true, "success"}
			wg.Done()
		case <-done:
			close(ch)
			return r, nil
		}
	}

}
