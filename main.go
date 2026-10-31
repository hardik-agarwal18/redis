package main

import (
	"fmt"
)

func main() {
	s := NewStore()
	s.Set("a", "42")
	s.Set("b", "52")

	val, _ := s.Get("a")

	fmt.Println("Hello")
	fmt.Println(s.data)
	s.Delete("b")

	fmt.Println(s.data)
	fmt.Println(val)
}
