package tests

import (
	"reflect"
	"strconv"
	"testing"
)

/**
 * [728] Self Dividing Numbers
 *
 *
 * A self-dividing number is a number that is divisible by every digit it contains.
 *
 * For example, 128 is a self-dividing number because 128 % 1 == 0, 128 % 2 == 0, and 128 % 8 == 0.
 *
 * Also, a self-dividing number is not allowed to contain the digit zero.
 *
 * Given a lower and upper number bound, output a list of every possible self dividing number, including the bounds if possible.
 *
 * Example 1:
 *
 * Input:
 * left = 1, right = 22
 * Output: [1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 15, 22]
 *
 *
 *
 * Note:
 * The boundaries of each input argument are 1 <= left <= right <= 10000.
 *
 */

func TestSelfDividingNumbers(t *testing.T) {
	var cases = []struct {
		left   int
		right  int
		output []int
	}{
		{
			left:   1,
			right:  22,
			output: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 15, 22},
		},
	}
	for _, c := range cases {
		x := selfDividingNumbers(c.left, c.right)
		if !reflect.DeepEqual(x, c.output) {
			t.Fail()
		}
	}
}

// submission codes start here

func selfDividingNumbers(left int, right int) []int {
	res := []int{}
	for i := left; i <= right; i++ {
		var flag bool
		str := strconv.Itoa(i)
		for _, v := range str {
			x := int(v-'0')
			if x != 0 {
				if i % x != 0 {
					flag = true
					break
				}
			} else {
				flag = true
				break
			}
		}
		if !flag {
			res = append(res, i)
		}
	}
	return res
}

// submission codes end
