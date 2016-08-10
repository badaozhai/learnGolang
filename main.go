package main

import (
	"fmt"
	"learnGolang/day_03"
	"reflect"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	var sli []interface{}
	sli = append(sli, 3)
	sli = append(sli, "abc")
	sli = append(sli, nil)
	sli = append(sli, "abcd")
	sli = append(sli, Person{"张三", 18})

	p := Person{"张三", 18}
	index := indexOfSlince(p, sli)

	fmt.Println("index", index)

	var sli2 []string
	sli2 = append(sli2, "abc")
	sli2 = append(sli2, "ddd")

	//http://stackoverflow.com/questions/12990338/cannot-convert-string-to-interface
	//官方解释
	//https://golang.org/doc/faq#convert_slice_of_interface
	//Can I convert a []T to an []interface{}?
	//
	//Not directly, because they do not have the same representation in memory. It is necessary to copy the elements individually to the destination slice. This example converts a slice of int to a slice of interface{}:
	//
	//t := []int{1, 2, 3, 4}
	//s := make([]interface{}, len(t))
	//for i, v := range t {
	//s[i] = v
	//}

	//why??? 怎么才可以让这里通用?
	new := make([]interface{}, len(sli2))
	for i, v := range sli2 {
		new[i] = v
	}
	index2 := indexOfSlince("abc3", new)
	fmt.Println(index2)
}

func indexOfSlince(item interface{}, sli []interface{}) int {
	if len(sli) == 0 {
		//len长度为0
		return -1
	}
	//typeOfItem:= reflect.TypeOf(item)
	//fmt.Println(typeOf)
	//fmt.Println(reflect.TypeOf(typeOf))
	for index, value := range sli {
		if item == value {
			return index
		}
		fmt.Println(index, value, reflect.TypeOf(value))
	}
	return -1
}
