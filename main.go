package main

import "fmt"

type Store struct {
	data map[string]string
}

func (s *Store) Get(key string) (string, bool) {
	val, ok := s.data[key]

	return val, ok
}

func (s *Store) Set(key string, value string) {
	s.data[key] = value
}

func (s *Store) Delete(key string) {
	delete(s.data, key)
}

func NewStore() *Store {
	return &Store{
		data: make(map[string]string),
	}
}

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
