package main

import "fmt"

func main() {
	s := make([]string, 3)
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println(s)

	s = append(s, "d", "e")

	fmt.Println(s)

	c := make([]string, len(s))
	copy(c, s)

	fmt.Println(c)

	fmt.Println(c[2:4])
	fmt.Println(c[2:])
	fmt.Println(c[:3])

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		twoD[i] = make([]int, 3);
		for j := 0; j < 3; j++ {
			twoD[i][j] = i+j
		}
	}

	fmt.Println(twoD)
}
