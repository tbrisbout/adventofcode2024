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

func parseInput(input string) (res [2][]int) {
	for _, s := range strings.Split(input, "\n") {
		a, b, _ := strings.Cut(s, "   ")
		res[0] = append(res[0], toInt(a))
		res[1] = append(res[1], toInt(b))
	}
	return res
}

func SumDistance(input string) (sum int) {
	res := parseInput(input)

	slices.Sort(res[0])
	slices.Sort(res[1])

	for i := range res[0] {
		sum += int(math.Abs(float64(res[0][i] - res[1][i])))
	}

	return sum
}

func SimilarityScore(input string) (score int) {
	res := parseInput(input)

	for _, a := range res[0] {
		count := 0
		for _, b := range res[1] {
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
