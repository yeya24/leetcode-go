package tests

import (
	"reflect"
	"testing"
)

/**
 * [338] Counting Bits
 *
 * Given a non negative integer number num. For every numbers i in the range 0 <= i <= num calculate the number of 1's in their binary representation and return them as an array.
 *
 * Example 1:
 *
 *
 * Input: 2
 * Output: [0,1,1]
 *
 * Example 2:
 *
 *
 * Input: 5
 * Output: [0,1,1,2,1,2]
 *
 *
 * Follow up:
 *
 *
 * 	It is very easy to come up with a solution with run time O(n*sizeof(integer)). But can you do it in linear time O(n) /possibly in a single pass?
 * 	Space complexity should be O(n).
 * 	Can you do it like a boss? Do it without using any builtin function like __builtin_popcount in c++ or in any other language.
 *
 */

func TestCountingBits(t *testing.T) {
	var cases = []struct {
		input  int
		output []int
	}{
		{
			input:  2,
			output: []int{0, 1, 1},
		},
		{
			input:  5,
			output: []int{0, 1, 1, 2, 1, 2},
		},
	}
	for _, c := range cases {
		x := countBits(c.input)
		if !reflect.DeepEqual(x, c.output) {
			t.Fail()
		}
	}
}

// submission codes start here

func countBits(num int) []int {
	if num == 0 {
		return []int{0}
	}
	if num == 1 {
		return []int{0, 1}
	}
	res := make([]int, num+1)
	res[0], res[1] = 0, 1
	for i := 2; i <= num; i++ {
	    x := i
	    for x > 0 {
	        res[i] += x & 1
	        x >>= 1
        }
    }
	return res
}

// submission codes end
