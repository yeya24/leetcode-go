package tests

import (
	"reflect"
	"testing"
)

/**
 * [1195] Distribute Candies to People
 *
 * We distribute some number of candies, to a row of n = num_people people in the following way:
 *
 * We then give 1 candy to the first person, 2 candies to the second person, and so on until we give n candies to the last person.
 *
 * Then, we go back to the start of the row, giving n + 1 candies to the first person, n + 2 candies to the second person, and so on until we give 2 * n candies to the last person.
 *
 * This process repeats (with us giving one more candy each time, and moving to the start of the row after we reach the end) until we run out of candies.  The last person will receive all of our remaining candies (not necessarily one more than the previous gift).
 *
 * Return an array (of length num_people and sum candies) that represents the final distribution of candies.
 *
 *
 * Example 1:
 *
 *
 * Input: candies = 7, num_people = 4
 * Output: [1,2,3,1]
 * Explanation:
 * On the first turn, ans[0] += 1, and the array is [1,0,0,0].
 * On the second turn, ans[1] += 2, and the array is [1,2,0,0].
 * On the third turn, ans[2] += 3, and the array is [1,2,3,0].
 * On the fourth turn, ans[3] += 1 (because there is only one candy left), and the final array is [1,2,3,1].
 *
 *
 * Example 2:
 *
 *
 * Input: candies = 10, num_people = 3
 * Output: [5,2,3]
 * Explanation:
 * On the first turn, ans[0] += 1, and the array is [1,0,0].
 * On the second turn, ans[1] += 2, and the array is [1,2,0].
 * On the third turn, ans[2] += 3, and the array is [1,2,3].
 * On the fourth turn, ans[0] += 4, and the final array is [5,2,3].
 *
 *
 *
 * Constraints:
 *
 *
 * 	1 <= candies <= 10^9
 * 	1 <= num_people <= 1000
 *
 *
 */

func TestDistributeCandiestoPeople(t *testing.T) {
	var cases = []struct {
		candies int
		people  int
		output  []int
	}{
		{
			candies: 7,
			people:  4,
			output:  []int{1, 2, 3, 1},
		},
		{
			candies: 10,
			people:  3,
			output:  []int{5, 2, 3},
		},
	}
	for _, c := range cases {
		x := distributeCandiesToPeople(c.candies, c.people)
		if !reflect.DeepEqual(x, c.output) {
			t.Fail()
		}
	}
}

// submission codes start here

func distributeCandiesToPeople(candies int, num_people int) []int {
	res := make([]int, num_people)
	for i := 0; candies > 0; candies -= i {
		x := i + 1
		res[i%num_people] += min(candies, x)
		i++
	}
	return res
}

// slow brute force
func distributeCandiesToPeople2(candies int, num_people int) []int {
	res := make([]int, num_people)
	c := 0
	for candies > 0 {
		for i := 0; i < num_people; i++ {
			c++
			if candies <= c {
				res[i] += candies
				return res
			}
			res[i] += c
			candies -= c
		}
	}
	return res
}

// submission codes end
