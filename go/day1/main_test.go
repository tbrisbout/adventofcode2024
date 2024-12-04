package main

import (
	"reflect"
	"strings"
	"testing"
)

var testInput = strings.TrimSpace(`
3   4
4   3
2   5
1   3
3   9
3   3
`)

func TestParseInput(t *testing.T) {
	got := parseInput(testInput)
	want := [2][]int{{3, 4, 2, 1, 3, 3}, {4, 3, 5, 3, 9, 3}}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("expected %+v but got %+v", want, got)
	}
}

func TestSumDistance(t *testing.T) {
	got := SumDistance(testInput)
	want := 11

	if got != want {
		t.Fatalf("expected %d but got %d", want, got)
	}
}

func TestSimilarityScore(t *testing.T) {
	got := SimilarityScore(testInput)
	want := 31

	if got != want {
		t.Fatalf("expected %d but got %d", want, got)
	}
}
