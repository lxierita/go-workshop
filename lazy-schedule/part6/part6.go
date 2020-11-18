package part6

import (
	"errors"
	"sync"
)

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

//Job describes the execution result of each job
type Job struct {
	Result  int
	Ok      bool
	Message string
}
type jobs struct {
	index  int
	result int
}

//Run scheduled functions in the order that they were added
func (s *Scheduler) Run() (r []Job, err []error) {
	var wg sync.WaitGroup
	wg.Add(len(s.scheduled))

	r = make([]Job, len(s.scheduled))
	var timeout, job chan jobs
	done := make(chan bool)

	for i, f := range s.scheduled {
		go func(i int, f intfunc) {
			job <- jobs{i, f()}
			if last := len(s.scheduled) - 1; i == last {
				wg.Wait()
				defer close(done)
				done <- true
			}
		}(i, f)
	}
	for {
		select {
		case failure := <-timeout:
			err = append(err, errors.New("timed out"))
			r[failure.index] = Job{0, false, "timed out"}
			wg.Done()
		case success := <-job:
			r[success.index] = Job{success.result, true, "success"}
			wg.Done()
		case <-done:
			close(timeout)
			close(job)
			return r, nil
		}
	}

}
