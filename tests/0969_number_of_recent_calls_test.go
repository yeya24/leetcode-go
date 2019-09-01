package tests

import (
	"testing"
)

/**
 * [969] Number of Recent Calls
 *
 * Write a class RecentCounter to count recent requests.
 *
 * It has only one method: ping(int t), where t represents some time in milliseconds.
 *
 * Return the number of pings that have been made from 3000 milliseconds ago until now.
 *
 * Any ping with time in [t - 3000, t] will count, including the current ping.
 *
 * It is guaranteed that every call to ping uses a strictly larger value of t than before.
 *
 *
 *
 * Example 1:
 *
 *
 * Input: inputs = [[],[1],[100],[3001],[3002]]
 * Output: [null,1,2,3,3]
 *
 *
 *
 * Note:
 *
 *
 * 	Each test case will have at most 10000 calls to ping.
 * 	Each test case will call ping with strictly increasing values of t.
 * 	Each call to ping will have 1 <= t <= 10^9.
 *
 *
 *
 *
 *
 */

func TestNumberofRecentCalls(t *testing.T) {

}

// submission codes start here

type RecentCounter struct {
	times []int
	count int
}

func RecentCounterConstructor() RecentCounter {
	return RecentCounter{times: make([]int, 0), count: 0}
}

func (this *RecentCounter) Ping(t int) int {
	this.times = append(this.times, t)
	for {
		if t >= 3000 && t - 3000 > this.times[this.count] {
			this.count++
		} else {
			break
		}
	}
	return len(this.times) - this.count
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */

// submission codes end
