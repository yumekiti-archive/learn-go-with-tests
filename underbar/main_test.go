package main

import (
	"testing"
)

func TestMain(t *testing.T) {
	t.Run("loop", func(t *testing.T) {
		got := Loop([]int{20, 30, 40})
		want := "[20 30 40]"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
