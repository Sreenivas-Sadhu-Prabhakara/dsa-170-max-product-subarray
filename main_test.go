package main

import "testing"

func TestMaxProduct(t *testing.T) {
	cases := []struct {
		in   []int
		want int
	}{
		{[]int{2, 3, -2, 4}, 6},
		{[]int{-2, 0, -1}, 0},
		{[]int{-2, 3, -4}, 24}, // the two negatives multiply through
		{[]int{-5}, -5},
		{[]int{0, 2}, 2},
		{[]int{-1, -2, -3}, 6},
		{[]int{2, -5, 1, 0, 3}, 3},
		{[]int{}, 0},
	}
	for _, c := range cases {
		if got := MaxProduct(c.in); got != c.want {
			t.Errorf("MaxProduct(%v) = %d, want %d", c.in, got, c.want)
		}
	}
}
