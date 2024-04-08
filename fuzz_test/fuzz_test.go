package main

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func FuzzSum(f *testing.F) {
	f.Add(10)
	f.Fuzz(func(t *testing.T, n int) {
		n %= 20

		var vals []int64
		var expect int64
		for i := 0; i < n; i++ {
			val := rand.Int63() % 1e6
			vals = append(vals, val)
			expect += val
		}

		assert.Equal(t, expect, Sum(vals))
	})
}
