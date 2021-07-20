package main

import (
	"fmt"
	"math/rand"
)

type Employee interface {
	calsal(timespan int) int
}

type Fulltime struct {
	basic int
}
type Contactor struct {
	basic int
}
type Freelancer struct {
	basic int
}

func (e Fulltime) calsal(days int) int {
	return days * e.basic
}

func (e Contactor) calsal(days int) int {
	return days * e.basic
}

func (e Freelancer) calsal(hours int) int {
	if hours < 20 {
		return 0
	}
	return hours * e.basic
}


func main(){
	f := Fulltime{basic:500}
	c := Contactor{basic: 100}
	fr := Freelancer{basic: 10}
	data := [3]Employee{f, c, fr}
	for idx, elem := range data {
		if idx == 2{
			fmt.Println(elem.calsal(rand.Intn(200)))
		}else {
			fmt.Println(elem.calsal(rand.Intn(28)))
		}
	}
}