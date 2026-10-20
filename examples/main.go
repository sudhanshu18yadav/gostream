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
	nums := stream.Of(1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20, 21, 22, 23, 24, 25, 26, 27, 28, 29)
	f1 := func(value int) string {
		switch {
			case value <= 10 : return "ten"
			case  value <= 20 : return "twenty"
			default : return "thirty"
		}
	}

	rst := stream.GroupBy(nums, f1)
// 	rst := stream.OfMap(stream.GroupBy(nums, f1)).Map(func(value item.Pair[string, []int]) item.Pair[string, []int] {
// 		return
// })

	fmt.Println(rst)
}
