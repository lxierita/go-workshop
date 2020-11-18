package part4

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

	r := make([]int, len(s.scheduled))

	ch := make(chan jobs, 10000)

	for i, f := range s.scheduled {
		go func(i int, f intfunc) {
			ch <- jobs{i, f()}
		}(i, f)
	}
	for i := 0; i < len(s.scheduled); i++ {
		select {
		case j := <-ch:
			r[j.index] = j.result
		}
	}
	return r
}
