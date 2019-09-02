package tests

import (
	"testing"
)

/**
 * [1280] Diet Plan Performance
 *
 * A dieter consumes calories[i] calories on the i-th day.  For every consecutive sequence of k days, they look at T, the total calories consumed during that sequence of k days:
 *
 * 	If T < lower, they performed poorly on their diet and lose 1 point;
 * 	If T > upper, they performed well on their diet and gain 1 point;
 * 	Otherwise, they performed normally and there is no change in points.
 *
 * Return the total number of points the dieter has after all calories.length days.
 * Note that: The total points could be negative.
 *
 * Example 1:
 *
 * Input: calories = [1,2,3,4,5], k = 1, lower = 3, upper = 3
 * Output: 0
 * Explaination: calories[0], calories[1] < lower and calories[3], calories[4] > upper, total points = 0.
 * Example 2:
 *
 * Input: calories = [3,2], k = 2, lower = 0, upper = 1
 * Output: 1
 * Explaination: calories[0] + calories[1] > upper, total points = 1.
 *
 * Example 3:
 *
 * Input: calories = [6,5,0,0], k = 2, lower = 1, upper = 5
 * Output: 0
 * Explaination: calories[0] + calories[1] > upper, calories[2] + calories[3] < lower, total points = 0.
 *
 *
 * Constraints:
 *
 * 	1 <= k <= calories.length <= 10^5
 * 	0 <= calories[i] <= 20000
 * 	0 <= lower <= upper
 *
 */

func TestDietPlanPerformance(t *testing.T) {
	var cases = []struct {
		input  []int
		k      int
		lower  int
		upper  int
		output int
	}{
		{
			input:  []int{1, 2, 3, 4, 5},
			k:      1,
			lower:  3,
			upper:  3,
			output: 0,
		},
		{
			input:  []int{3, 2},
			k:      2,
			lower:  0,
			upper:  1,
			output: 1,
		},
		{
			input:  []int{6, 5, 0, 0},
			k:      2,
			lower:  1,
			upper:  5,
			output: 0,
		},
		{
			input:  []int{6, 13, 8, 7, 10, 1, 12, 11},
			k:      6,
			lower:  5,
			upper:  37,
			output: 3,
		},
	}
	for _, c := range cases {
		x := dietPlanPerformance(c.input, c.k, c.lower, c.upper)
		if x != c.output {
			t.Fail()
		}
	}
}

// submission codes start here

func dietPlanPerformance(calories []int, k int, lower int, upper int) int {
	l := len(calories)
	p := 0
	sum := 0
	for i := 0; i <= l-k; i++ {
		if i == 0 {
			for j := 0; j < k; j++ {
				sum += calories[i+j]
			}
		} else {
			sum -= calories[i-1]
			sum += calories[i+k-1]
		}
		if sum > upper {
			p += 1
		} else if sum < lower {
			p -= 1
		}
	}
	return p
}

// submission codes end
