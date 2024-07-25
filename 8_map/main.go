package main

import (
	"fmt"
	"maps"
)

func main() {
	// var m = make(map[string]string)
	// m["name"] = "unknown"
	// m["who"] = "i"
	// fmt.Println(m)

	// var m = map[string]string{"name": "unknown", "id": "15a7"}

	// fmt.Println(m)
	// delete(m, "name")
	// fmt.Println(m)
	// clear(m)
	// fmt.Println(m)

	// var m = map[string]int{"name": 1, "id": 15}

	// delete(m, "name")

	// v,ok := m["name"]
	// if ok {

	// 	fmt.Println("ok", v)
	// } else {
	// 	fmt.Println("not ok", v)
	// }

	var m1 = map[string]int{"name": 1, "id": 15}
	var m2 = map[string]int{"name": 1, "id": 15}

	fmt.Println(maps.Equal(m1,m2))
}