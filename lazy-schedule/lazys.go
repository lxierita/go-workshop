package lazys

import (
	"reflect"
	"regexp"
	"runtime"
)

type Scheduler struct {
}

var Scheduled = make([]interface{}, 0)

func (s Scheduler) Add(f func(...int) int, args []int) {
	toAdd := func() int {
		return f(args...)
	}
	Scheduled = append(Scheduled, toAdd)
}

func (s Scheduler) Run() []int {
	var result []int
	for _, f := range Scheduled {
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
