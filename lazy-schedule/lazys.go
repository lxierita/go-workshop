package lazys

import (
	"reflect"
	"regexp"
	"runtime"
)

type Scheduler struct {
	Scheduled []interface{}
}

func (s Scheduler) Add(f func(...int) int, args []int) int {
	toAdd := func() int {
		return f(args...)
	}
	s.Scheduled = append(s.Scheduled, toAdd)
	return len(s.Scheduled)
}

func (s Scheduler) Run() []int {
	var result []int

	for _, f := range s.Scheduled {
		result = append(result, f.(func() int)())
	}
	return result
}

func GetName(f func(...int) int) string {
	var s, name string

	s = runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
	regex := regexp.MustCompile("[^\x2E]*$")
	name = regex.FindString(s)
	return name
}
