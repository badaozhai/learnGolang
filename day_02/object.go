package day_02

import "fmt"

type Rectangle struct {
	height float64
	witdh  float64
}
type Circle struct {
	radius float64
}

/**
这仅仅是一个函数
*/
func area1(r Rectangle) float64 {
	return r.height * r.witdh
}

/*
接收者是 矩形 类似于java里的方法 附属于对象,实例方法
*/
func (r Rectangle) area() float64 {
	return r.height * r.witdh
}

/*
接收者是圆形,类似于java中的方法,Circle里的一个方法,
*/
func (c Circle) area() float64 {
	return c.radius * c.radius * 3.14
}
func Main() {
	r := Rectangle{8, 6}
	fmt.Println("")
	fmt.Println(area1(r))
	fmt.Println(r.area())
	c := Circle{3}
	fmt.Println("C的面积", c.area())
}
