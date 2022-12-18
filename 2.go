package main

import "fmt"

type animals interface {
	yell()
}
type pets struct {
	Name  string
	voice string
}

func (a *pets) yell() {
	fmt.Println(a.Name + " 叫声是：" + a.voice)
}

func main() {
	var a animals = &pets{Name: "Dog", voice: "wangwang"}
	var b animals = &pets{Name: "Cat", voice: "mewmew"}
	a.yell()
	b.yell()
}
