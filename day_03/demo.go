package day_03

import "fmt"

type Saiyan struct {
	Name  string
	Power int
}

func Main() {
	goku := Saiyan{"Goku", 9000}
	Super(goku)
	fmt.Println(goku.Power)
}

func Super(s Saiyan) {
	s.Power += 10000
}
