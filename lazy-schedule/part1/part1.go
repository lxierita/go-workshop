package part1

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

//Add functions to Scheduler lazily
func (s *Scheduler) Add(f intfunc, args []int) {
	af := func(...int) int {
		return f(args...)
	}
	s.scheduled = append(s.scheduled, af)
}
