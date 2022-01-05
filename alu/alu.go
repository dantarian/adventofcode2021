package alu

import (
	"fmt"
	"math"
)

func MaxValidFinder() int64 {
	zShrink := []bool{false, false, false, true, false, false, true, true, true, false, false, true, true, true}
	xAdd := []int64{11, 13, 15, -8, 13, 15, -11, -4, -15, 14, 14, -1, -8, -14}
	wAdd := []int64{6, 14, 14, 10, 9, 12, 8, 13, 12, 6, 9, 15, 4, 10}

	_, value := validDigit(0, zShrink, xAdd, wAdd, 0, Descending)
	return value
}

func MinValidFinder() int64 {
	zShrink := []bool{false, false, false, true, false, false, true, true, true, false, false, true, true, true}
	xAdd := []int64{11, 13, 15, -8, 13, 15, -11, -4, -15, 14, 14, -1, -8, -14}
	wAdd := []int64{6, 14, 14, 10, 9, 12, 8, 13, 12, 6, 9, 15, 4, 10}

	_, value := validDigit(0, zShrink, xAdd, wAdd, 0, Ascending)
	return value
}

type direction int

const (
	Ascending direction = iota
	Descending
)

func validDigit(z int64, zShrink []bool, xAdd []int64, wAdd []int64, prefix int64, d direction) (bool, int64) {
	startDigit, endDigit, step := int64(1), int64(9), int64(1)
	condition := func(val, target int64) bool { return val <= target }
	if d == Descending {
		startDigit, endDigit, step = int64(9), int64(1), int64(-1)
		condition = func(val, target int64) bool { return val >= target }
	}
	for candidate := startDigit; condition(candidate, endDigit); candidate += step {
		newPrefix := prefix*10 + candidate
		if len(zShrink) == 10 {
			fmt.Printf("%v,000,000,000\n", newPrefix)
		}

		w, z := candidate, z
		x := (z % 26) + xAdd[0]
		if zShrink[0] {
			z = z / 26
		}
		if w != x {
			z = z*26 + w + wAdd[0]
		}

		if z > int64(math.Pow(26, float64(len(zShrink)-1))) {
			continue
		}

		if len(zShrink) == 1 {
			if z != 0 {
				return false, candidate
			}
			return true, candidate
		}

		valid, value := validDigit(z, zShrink[1:], xAdd[1:], wAdd[1:], newPrefix, d)
		if !valid {
			continue
		}
		return true, candidate*int64(math.Pow10(len(zShrink)-1)) + value
	}
	return false, 0
}
