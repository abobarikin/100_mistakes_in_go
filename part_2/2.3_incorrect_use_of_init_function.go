package main

import "fmt"

var a = func() int {
	fmt.Println("var")
	return 0
}()

func init() {
	fmt.Println("init")
}

func main2() {
	fmt.Println("main")
}