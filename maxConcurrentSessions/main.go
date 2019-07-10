package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)

type Line struct {
	start int
	end   int
}

// sorte interface
type Lines []Line
type MaxRanges []Lines

func (a Lines) Len() int           { return len(a) }
func (a Lines) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Lines) Less(i, j int) bool { return a[i].start < a[j].start }

// Método: função com receiver (receptor)
func (l1 Line) GetOverlap(l2 Line) int {
	ol := l1.getShiftOverlap(l2)
	if ol == 0 {
		ol = l1.getContainedOverlap(l2)
	}
	return ol
}

// Método: função com receiver (receptor)
func (l1 Line) Length() int {
	return l1.end - l1.start
}

func (l1 Line) getContainedOverlap(l2 Line) int {
	smaller := l1
	bigger := l2
	if smaller.Length() > bigger.Length() {
		smaller, bigger = bigger, smaller
	}
	if smaller.start >= bigger.start && smaller.start <= bigger.end && smaller.end >= bigger.start && smaller.end <= bigger.end {
		return smaller.Length()
	}

	return 0
}

func (l1 Line) getShiftOverlap(l2 Line) int {
	first := l1
	second := l2
	if first.start > second.start {
		first, second = l2, l1
	}

	// fmt.Println("first, second", first, second)
	if first.end > second.start && first.end <= second.end {
		return first.end - second.start
	}
	return 0
}

func (lines Lines) sortByStart() {
	sort.Sort(Lines(lines))
}

func (a Line) Equals(b Line) bool {
	if a.end == b.end && a.start == b.start {
		return true
	}
	return false
}

// Equal tells whether a and b contain the same elements.
// A nil argument is equivalent to an empty slice.
func (a Lines) Equals(b Lines) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if !v.Equals(b[i]) {
			return false
		}
	}
	return true
}

func (a MaxRanges) toString() string {
	result := "( from "

	distinctLines := Lines{}
	for i, arr := range a {
		min := math.MaxInt64
		max := 0

		found := false
		for _, v := range arr {
			if v.start < min {
				min = v.start
				found = distinctLines.addIfNotExists(v)
			}
			if v.end > max {
				max = v.end
				found = distinctLines.addIfNotExists(v)
			}
		}
		if found {
			strRange := strconv.Itoa(min) + " to " + strconv.Itoa(max)

			if i > 0 {
				result += " or "
			}
			result += strRange
		}
	}
	result += " )"

	return result
}

func (lines *Lines) addIfNotExists(line Line) bool {
	if !lines.Contains(line) {
		*lines = append(*lines, line)
		return true
	}

	return false
}

func (lines Lines) Contains(line Line) bool {
	for _, l := range lines {
		if l.Equals(line) {
			return true
		}
	}
	return false
}

func evaluateOverlap(lines Lines) (max int, maxRange MaxRanges) {
	max = 0
	maxRange = MaxRanges{}

	for i, curr := range lines {
		currOverlap := 0
		rng := Lines{curr}

		for j, next := range lines {
			if i == j {
				continue
			}
			overlap := curr.GetOverlap(next)

			if overlap > 0 {
				currOverlap += overlap

				rng = append(rng, next)
			}
		}

		if len(rng) > 1 {
			if currOverlap > max {
				max = currOverlap
				maxRange = MaxRanges{}
				maxRange = append(maxRange, rng)
			} else {
				maxRange = append(maxRange, rng)
			}
		}
	}

	return
}

/*
https://www.careercup.com/question?id=5743263735087104

Find the maximum number of concurrent sessions in the following data with the first value representing start time and last value end time. The input is not necessarily sorted.
Input: (2,5), (3,6), (8,10),(10,12),(9,20)
Output: 3 (from 8 to 20)
Input: (2,5), (3,6), (8,10),(9,12),(12,20)
Output: 2 (from 8 to 12 or 2 to 6)
*/
func main() {

	lines1 := Lines{Line{2, 5}, Line{3, 6}, Line{8, 10}, Line{10, 12}, Line{9, 20}}
	lines2 := Lines{Line{2, 5}, Line{3, 6}, Line{8, 10}, Line{9, 12}, Line{12, 20}}
	// lines.sortByStart()

	max, maxRange := evaluateOverlap(lines1)
	fmt.Println("Max", max, maxRange.toString())

	max, maxRange = evaluateOverlap(lines2)
	fmt.Println("Max", max, maxRange.toString())

}
