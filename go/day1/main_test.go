package main

import (
	"strings"
	"testing"
)

func TestSumDistance(t *testing.T) {
	testInput := strings.TrimSpace(`
3   4
4   3
2   5
1   3
3   9
3   3
`)

	got := SumDistance(testInput)
	want := 11

	if got != want {
		t.Fatalf("expected %d but got %d", want, got)
	}

}
