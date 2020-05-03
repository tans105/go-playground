package main

import "fmt"

func main() {
	s := "i'm sorry dave i can't do that"
	fmt.Println(s)
	fmt.Println([]byte(s))//[105 39 109 32 115 111 114 114 121 32 100 97 118 101 32 105 32 99 97 110 39 116 32 100 111 32 116 104 97 116]
	fmt.Println(string([]byte(s))) //i'm sorry dave i can't do that
	fmt.Println(string([]byte(s)[:14])) //i'm sorry dave
	fmt.Println(string([]byte(s)[10:22])) //dave i can't
	fmt.Println(string([]byte(s)[17:])) //can't do that

	for _, v := range []byte(s) {
		fmt.Print(string(v) + "\t") //i	'	m	 	s	o	r	r	y	 	d	a	v	e	 	i	 	c	a	n	'	t	 	d	o	 	t	h	a	t
	}
}
