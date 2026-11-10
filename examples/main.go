package main

import (
	"fmt"
	"reflect"
	"runtime"

	"github.com/mariomac/gostream/stream"
)

func GetFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func main() {
	nums := stream.Of(1,4,5,6,7,2, 1,3,8)
	fmt.Println(nums.ToSlice())
	lessThan5 := func(val int) bool{
		return val < 5
	}
	// t1 := nums.TakeWhile(lessThan5)
	t1 := nums.Filter(lessThan5)

	fmt.Println(t1.ToSlice())
}
