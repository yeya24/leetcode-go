package tests

import (
	"testing"
)

/**
 * [1287] Distance Between Bus Stops
 *
 * A bus has n stops numbered from 0 to n - 1 that form a circle. We know the distance between all pairs of neighboring stops where distance[i] is the distance between the stops number i and (i + 1) % n.
 *
 * The bus goes along both directions i.e. clockwise and counterclockwise.
 *
 * Return the shortest distance between the given start and destination stops.
 *
 *
 * Example 1:
 *
 *
 *
 *
 * Input: distance = [1,2,3,4], start = 0, destination = 1
 * Output: 1
 * Explanation: Distance between 0 and 1 is 1 or 9, minimum is 1.
 *
 *
 *
 * Example 2:
 *
 *
 *
 *
 * Input: distance = [1,2,3,4], start = 0, destination = 2
 * Output: 3
 * Explanation: Distance between 0 and 2 is 3 or 7, minimum is 3.
 *
 *
 *
 *
 * Example 3:
 *
 *
 *
 *
 * Input: distance = [1,2,3,4], start = 0, destination = 3
 * Output: 4
 * Explanation: Distance between 0 and 3 is 6 or 4, minimum is 4.
 *
 *
 *
 * Constraints:
 *
 *
 * 	1 <= n <= 10^4
 * 	distance.length == n
 * 	0 <= start, destination < n
 * 	0 <= distance[i] <= 10^4
 *
 */

func TestDistanceBetweenBusStops(t *testing.T) {
	var cases = []struct {
		input  []int
		start  int
		dest   int
		output int
	}{
		{
			input:  []int{1,2,3,4},
			start: 0,
			dest: 1,
			output: 1,
		},
		{
			input:  []int{1,2,3,4},
			start: 0,
			dest: 2,
			output: 3,
		},
		{
			input:  []int{1,2,3,4},
			start: 0,
			dest: 3,
			output: 4,
		},
	}
	for _, c := range cases {
		x := distanceBetweenBusStops(c.input, c.start, c.dest)
		if x != c.output {
			t.Fail()
		}
	}
}

// submission codes start here

func distanceBetweenBusStops(distance []int, start int, destination int) int {
	sum := 0
	if start > destination {
		start, destination = destination, start
	}
	for i := start; i < destination; i++ {
		sum += distance[i]
	}
	z := 0
	for i := 0; i < len(distance); i++ {
		z += distance[i]
	}
	return min(sum, z-sum)
}

// submission codes end
