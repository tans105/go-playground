package main

import "fmt"

type person struct {
	fname string
	lname string
}

type sa struct {
	person
	ltk bool
}

func (p person) pSpeak() {
	fmt.Println(p.fname, `says, "Wassup!"`)
}

func (s sa) saSpeak() {
	fmt.Println(s.fname, "has a license to kill:", s.ltk)
}

func main() {
	p := person{"Tanmay", "Awasthi"}
	s := sa{person{"James", "Bond"}, true}
	fmt.Println(p.fname) //Tanmay
	p.pSpeak()           //Tanmay says, "Wassup!"
	fmt.Println(s.fname) //James
	s.saSpeak()          //James has a license to kill: true
	s.pSpeak()           //James says, "Wassup!"
	s.person.pSpeak()	 //James says, "Wassup!"
}
