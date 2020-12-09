package _001_Two_Sum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type QnA struct {
	nums   []int
	target int
	result []int
}

func TestTwoSum(t *testing.T) {
	qas := []QnA{
		{nums: []int{2, 7, 11, 15}, target: 9, result: []int{0, 1}},
		{nums: []int{3, 2, 4}, target: 6, result: []int{1, 2}},
		{nums: []int{3, 3}, target: 6, result: []int{0, 1}},
		{nums: []int{3, 2, 8, 4}, target: 11, result: []int{0, 2}},
	}

	for _, v := range qas {
		assert.Equal(t, v.result, twoSum(v.nums, v.target))
	}
}
