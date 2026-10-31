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

func TestKeys_EmptyStore(t *testing.T) {
	store := NewStore()

	got := store.Keys()

	if len(got) != 0 {
		t.Errorf("Keys() on empty store =%v, want empty slice", got)
	}
}

func TestSetGet_RoundTrip(t *testing.T) {
	store := NewStore()

	store.Set("Hardik", "5")

	if val, ok := store.Get("Hardik"); val != "5" || !ok {
		t.Errorf("Get() Failed")
	}

	if _, ok := store.Get("Missing"); ok {
		t.Errorf("Get() Failed")
	}

}

// Table tests
func TestDelete(t *testing.T) {
	tests := []struct {
		name      string
		setup     map[string]string
		deleteKey string
		wantLen   int
	}{
		{"deletes existing key", map[string]string{"a": "1", "b": "2"}, "a", 1},
		{"no-op on missing key", map[string]string{"a": "1", "b": "2"}, "c", 2},
		{"no-op on empty store", map[string]string{}, "x", 0},
		{"deletes last remaining key", map[string]string{"x": "1"}, "x", 0},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			store := NewStore()

			for k, v := range tc.setup {
				store.Set(k, v)
			}

			store.Delete(tc.deleteKey)

			if got := store.Len(); got != tc.wantLen {
				t.Errorf("after Delete(%q) :Len()=%d, Wanted=%d", tc.deleteKey, got, tc.wantLen)
			}

			if _, ok := store.Get(tc.deleteKey); ok {
				t.Errorf("after Delete(%q): Key still present", tc.deleteKey)

			}
		})
	}

}
