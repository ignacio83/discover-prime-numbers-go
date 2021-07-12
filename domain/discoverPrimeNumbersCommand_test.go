package domain

import (
	"reflect"
	"testing"
)

func TestExecute(t *testing.T) {
	t.Run("Should discover 10 prime numbers", func(t *testing.T) {
		command := DiscoverPrimeNumbersCommand{Qty: 10}
		got := command.Execute()
		want := []int{1, 2, 3, 5, 7, 11, 13, 17, 19, 23}

		if !reflect.DeepEqual(got, want) {
			t.Fatalf("got %d want %d", got, want)
		}
	})

	t.Run("Should discover 1 prime numbers", func(t *testing.T) {
		command := DiscoverPrimeNumbersCommand{Qty: 1}
		got := command.Execute()
		want := []int{1}

		if !reflect.DeepEqual(got, want) {
			t.Fatalf("got %d want %d", got, want)
		}
	})
}
