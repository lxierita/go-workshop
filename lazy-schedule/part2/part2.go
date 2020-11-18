package part2

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
func (s *Scheduler) Run() (r []int) {
	for _, f := range s.scheduled {
		r = append(r, f())
	}
	return
}
