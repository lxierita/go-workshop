package lazys

import (
	"fmt"
	"reflect"
	"regexp"
	"runtime"
)

type Scheduler struct {
	Result []int
	Len    int
	fs     []interface{}
}

func (s Scheduler) Add(f func(...int) int, args []int) {
	toAdd := func() {
		f(args...)
	}
	s.fs = append(s.fs, toAdd)
	s.Len = len(s.fs)
}

func (s Scheduler) Run() {
	for i, f := range s.fs {
		fmt.Println(i, f)
	}
}

func GetName(f func(...int) int) string {
	var s, name string

	s = runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
	regex := regexp.MustCompile("[^\x2E]*$")
	name = regex.FindString(s)
	return name
}
