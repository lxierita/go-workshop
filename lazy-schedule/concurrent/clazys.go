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
	c := cal()
	for {
		select {
		case i := <-c:
			result = append(result, i)
		default:
			return result
		}
	}
}

func cal() <-chan int {
	c := make(chan int)

	defer close(c)

	//one routine for each
	for _, f := range Scheduled {
		go func() { c <- f.(func(...int) int)() }()
	}

	// divide task in half
	// h := len(Scheduled) / 2
	// go func() {
	// 	for _, f := range Scheduled[:h] {
	// 		c <- f.(func(...int) int)()
	// 	}
	// }()
	// go func() {
	// 	for _, f := range Scheduled[h:] {
	// 		c <- f.(func(...int) int)()
	// 	}
	// }()

	return c
}

func GetName(f func(...int) int) string {
	var s, name string

	s = runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
	regex := regexp.MustCompile("[^\x2E]*$")
	name = regex.FindString(s)
	return name
}
