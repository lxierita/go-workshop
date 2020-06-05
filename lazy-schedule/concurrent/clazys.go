package clazys

import (
	"reflect"
	"regexp"
	"runtime"
)

type Scheduler struct {
}

var Scheduled = make([]interface{}, 0)

func (s Scheduler) Add(f func(...int) int, args []int) int {
	toAdd := func() int {
		return f(args...)
	}
	Scheduled = append(Scheduled, toAdd)
	return len(Scheduled)
}

func (s Scheduler) Run() []int {
	var result []int

	for _, f := range Scheduled {
		go cal(result, f)
	}
	return result
}

func cal(r []int, f interface{}) {
	r = append(r, f.(func() int)())
}

func GetName(f func(...int) int) string {
	var s, name string

	s = runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
	regex := regexp.MustCompile("[^\x2E]*$")
	name = regex.FindString(s)
	return name
}
