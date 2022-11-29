package main

import "fmt"

type Sleeper interface {
	sleep()
}

type Dog struct {
	Name string
}

type Cat struct {
	Name string
}

func (d Dog) sleep() {
	fmt.Printf("Dog %s is sleeping\n", d.Name)
}

func (c Cat) sleep() {
	fmt.Printf("Cat %s is sleeping\n", c.Name)
}
func AnimalSleep(s Sleeper) {
	s.sleep()
}
func main() {
	var s Sleeper
	dog := Dog{Name: "xiaobai"}
	cat := Cat{Name: "hellkity"}
	s = dog
	AnimalSleep(s)
	s = cat
	AnimalSleep(s)

	sList := []Sleeper{Dog{Name: "hhhh"}, Cat{Name: "ooooo"}}
	for _, s := range sList {
		s.sleep()
	}
}
