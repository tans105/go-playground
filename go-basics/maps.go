package main

import (
	"fmt"
)

func main() {

	m := make(map[string]int)

	m["foo"] = 1
	m["bar"] = 2

	fmt.Println(m)      // map[bar:2 foo:1]
	fmt.Println(len(m)) // 2

	delete(m, "foo")

	fmt.Println(m) // map[bar:2]

	var val int = m["bar"]
	fmt.Println(val) // 0

	var secondReturn bool = true
	_, secondReturn = m["foo"] // _ is used to ignored the value returned from m["foo"], we are interested in the second option return value
	fmt.Println(secondReturn)  // false

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println(n) // map[bar:2 foo:1]

}
