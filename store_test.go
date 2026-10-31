package main

import (
	"reflect"
	"testing"
)

func TestKeys_ReturnsAllKeysSorted(t *testing.T) {
	store := NewStore()

	store.Set("Hardik", "3")
	store.Set("Agarwal", "30")
	store.Set("Qwerty", "300")

	got := store.Keys()
	want := []string{"Agarwal", "Hardik", "Qwerty"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Keys() = %v, want %v", got, want)
		t.Errorf("The keys are different!!!")
	}
}
