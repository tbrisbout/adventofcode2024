package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

func toInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func parseInput(input string) (left, right []int) {
	for _, s := range strings.Split(input, "\n") {
		a, b, _ := strings.Cut(s, "   ")
		left = append(left, toInt(a))
		right = append(right, toInt(b))
	}
	return left, right
}

func SumDistance(input string) (sum int) {
	left, right := parseInput(input)

	slices.Sort(left)
	slices.Sort(right)

	for i := range left {
		sum += int(math.Abs(float64(left[i] - right[i])))
	}

	return sum
}

func SimilarityScore(input string) (score int) {
	left, right := parseInput(input)

	for _, a := range left {
		count := 0
		for _, b := range right {
			if b == a {
				count++
			}
		}

		score += a * count
	}

	return score
}

func main() {
	fmt.Println("Part 1:", SumDistance(input))
	fmt.Println("Part 2:", SimilarityScore(input))
}
