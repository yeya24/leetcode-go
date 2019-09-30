package tests

import (
	"testing"
)

/**
 * [1184] Car Pooling
 *
 * You are driving a vehicle that has capacity empty seats initially available for passengers.  The vehicle only drives east (ie. it cannot turn around and drive west.)
 * Given a list of trips, trip[i] = [num_passengers, start_location, end_location] contains information about the i-th trip: the number of passengers that must be picked up, and the locations to pick them up and drop them off.  The locations are given as the number of kilometers due east from your vehicle's initial location.
 * Return true if and only if it is possible to pick up and drop off all passengers for all the given trips.
 *
 * Example 1:
 *
 * Input: trips = [[2,1,5],[3,3,7]], capacity = 4
 * Output: false
 *
 *
 * Example 2:
 *
 * Input: trips = [[2,1,5],[3,3,7]], capacity = 5
 * Output: true
 *
 *
 * Example 3:
 *
 * Input: trips = [[2,1,5],[3,5,7]], capacity = 3
 * Output: true
 *
 *
 * Example 4:
 *
 * Input: trips = [[3,2,7],[3,7,9],[8,3,9]], capacity = 11
 * Output: true
 *
 *
 *
 *
 *
 *
 *
 *
 *
 *
 *
 * Constraints:
 *
 * 	trips.length <= 1000
 * 	trips[i].length == 3
 * 	1 <= trips[i][0] <= 100
 * 	0 <= trips[i][1] < trips[i][2] <= 1000
 * 	1 <= capacity <= 100000
 *
 */

func TestCarPooling(t *testing.T) {
	var cases = []struct {
		input  [][]int
		capa   int
		output bool
	}{
		{
			input:  [][]int{{2, 1, 5}, {3, 3, 7}},
			capa:   4,
			output: false,
		},
		{
			input:  [][]int{{2, 1, 5}, {3, 3, 7}},
			capa:   5,
			output: true,
		},
		{
			input:  [][]int{{3,2,7},{3,7,9},{8,3,9}},
			capa:   11,
			output: true,
		},
		{
			input:  [][]int{{2, 1, 5}, {3, 5, 7}},
			capa:   3,
			output: true,
		},
	}
	for _, c := range cases {
		x := carPooling(c.input, c.capa)
		if x != c.output {
			t.Fail()
		}
	}
}

// submission codes start here

func carPooling(trips [][]int, capacity int) bool {
	stops := [1001]int{}
	for _, v := range trips {
		stops[v[1]] += v[0]
		stops[v[2]] -= v[0]
	}
	currPassengers := 0
	for _, v := range stops {
		if v != 0 {
			currPassengers += v
			if currPassengers > capacity {
				return false
			}
		}
	}
	return true
}

// submission codes end
