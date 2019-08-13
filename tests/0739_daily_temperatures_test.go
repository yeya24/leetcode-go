package tests

import (
	"container/list"
	"reflect"
	"testing"
)

/**
 * [739] Daily Temperatures
 *
 *
 * Given a list of daily temperatures T, return a list such that, for each day in the input, tells you how many days you would have to wait until a warmer temperature.  If there is no future day for which this is possible, put 0 instead.
 *
 * For example, given the list of temperatures T = [73, 74, 75, 71, 69, 72, 76, 73], your output should be [1, 1, 4, 2, 1, 1, 0, 0].
 *
 *
 * Note:
 * The length of temperatures will be in the range [1, 30000].
 * Each temperature will be an integer in the range [30, 100].
 *
 */

func TestDailyTemperatures(t *testing.T) {
	var cases = []struct {
		input  []int
		output []int
	}{
		{
			input:  []int{73, 74, 75, 71, 69, 72, 76, 73},
			output: []int{1, 1, 4, 2, 1, 1, 0, 0},
		},
	}
	for _, c := range cases {
		x := dailyTemperatures(c.input)
		if !reflect.DeepEqual(x, c.output) {
			t.Fail()
		}
	}
}

// submission codes start here

func dailyTemperatures(T []int) []int {
	temp := make([]int, 0)
	res := make([]int, len(T))
	for i, t := range T {
		if len(temp) != 0 {
			for len(temp) > 0 {
				l := len(temp) - 1
				idx := temp[l]
				if T[idx] < t {
					res[idx] = i - idx
					temp = temp[:l]
				} else {
					break
				}
			}
		}
		temp = append(temp, i)
	}
	return res
}

// this way is too slow :(. maybe using container.list is not a good idea.
func dailyTemperatures1(T []int) []int {
	l := list.New()
	res := make([]int, len(T))
	for i, t := range T {
		if l.Len() == 0 {
			l.PushBack(i)
		} else {
			c := 0
			e := l.Front()
			for c < l.Len() {
				if T[e.Value.(int)] < t {
					newE := e.Next()
					idx := l.Remove(e).(int)
					res[idx] = i - idx
					e = newE
				} else {
					e = e.Next()
					c++
				}
			}
			l.PushBack(i)
		}
	}
	return res
}

// submission codes end
