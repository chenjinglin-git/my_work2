package main

import "fmt"

type introduce_myself struct {
	Name string
	Sex string
	Age int
	Height float64
	Weight float64
	Hobby string

}

func main() {
	a:=introduce_myself{
		Name: "chen jinglin",
		Sex: "男",
		Age: 18,
		Height: 172.0,
		Weight: 128.5,
		Hobby: "打篮球",
	}
	fmt.Println(a)
}