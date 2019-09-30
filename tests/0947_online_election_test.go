package tests

import (
	"testing"
)

/**
 * [947] Online Election
 *
 * In an election, the i-th vote was cast for persons[i] at time times[i].
 *
 * Now, we would like to implement the following query function: TopVotedCandidate.q(int t) will return the number of the person that was leading the election at time t.
 *
 * Votes cast at time t will count towards our query.  In the case of a tie, the most recent vote (among tied candidates) wins.
 *
 *
 *
 *
 * Example 1:
 *
 *
 * Input: [[[0,1,1,0,0,1,0],[0,5,10,15,20,25,30]],[3],[12],[25],[15],[24],[8]]
 * Output: [null,0,1,1,0,0,1]
 * Explanation:
 * At time 3, the votes are [0], and 0 is leading.
 * At time 12, the votes are [0,1,1], and 1 is leading.
 * At time 25, the votes are [0,1,1,0,0,1], and 1 is leading (as ties go to the most recent vote.)
 * This continues for 3 more queries at time 15, 24, and 8.
 *
 *
 *
 *
 * Note:
 *
 *
 * 	1 <= persons.length = times.length <= 5000
 * 	0 <= persons[i] <= persons.length
 * 	times is a strictly increasing array with all elements in [0, 10^9].
 * 	TopVotedCandidate.q is called at most 10000 times per test case.
 * 	TopVotedCandidate.q(int t) is always called with t >= times[0].
 *
 *
 *
 */

func TestOnlineElection(t *testing.T) {
	var cases = []struct {
		persons []int
		times   []int
		q       []int
		a       []int
	}{
		//{
		//	persons: []int{0, 1, 1, 0, 0, 1, 0},
		//	times:   []int{0, 5, 10, 15, 20, 25, 30},
		//	q:       []int{3, 12, 25, 15, 24, 8},
		//	a:       []int{0, 1, 1, 0, 0, 1},
		//},
		{
			persons: []int{0, 1, 0, 1, 1},
			times:   []int{24, 29, 31, 76, 81},
			q:       []int{28, 24, 29, 77, 30, 25, 76, 75, 81, 80},
			a:       []int{0, 0, 1, 1, 1, 0, 1, 0, 1, 1},
		},
	}
	for _, c := range cases {
		vote := Constructor(c.persons, c.times)
		for i := 0; i < len(c.q); i++ {
			if vote.Q(c.q[i]) != c.a[i] {
				t.Fail()
			}
		}
	}
}

// submission codes start here

type TopVotedCandidate struct {
	cache map[int]int
	times []int
}

func Constructor(persons []int, times []int) TopVotedCandidate {
	c := TopVotedCandidate{times: times, cache: make(map[int]int)}
	n := len(times)
	leader := -1
	m := make(map[int]int)
	for i := 0; i < n; i++ {
		m[persons[i]] += 1
		if leader == -1 || m[persons[i]] >= m[leader] {
			leader = persons[i]
		}
		c.cache[times[i]] = leader
	}
	return c
}

func (this *TopVotedCandidate) Q(t int) int {
	i := 0
	j := len(this.times) - 1
	for i <= j {
		m := (i + j) / 2
		if this.times[m] == t {
			return this.cache[this.times[m]]
		}
		if this.times[m] > t {
			j = m - 1
		} else if this.times[m] < t {
			i = m + 1
		}
	}
	return this.cache[this.times[i-1]]
}

/**
 * Your TopVotedCandidate object will be instantiated and called as such:
 * obj := Constructor(persons, times);
 * param_1 := obj.Q(t);
 */

// submission codes end
