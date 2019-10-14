package tests

import (
	"testing"
)

/**
 * [621] Task Scheduler
 *
 * Given a char array representing tasks CPU need to do. It contains capital letters A to Z where different letters represent different tasks. Tasks could be done without original order. Each task could be done in one interval. For each interval, CPU could finish one task or just be idle.
 *
 * However, there is a non-negative cooling interval n that means between two same tasks, there must be at least n intervals that CPU are doing different tasks or just be idle.
 *
 * You need to return the least number of intervals the CPU will take to finish all the given tasks.
 *
 *
 *
 * Example:
 *
 *
 * Input: tasks = ["A","A","A","B","B","B"], n = 2
 * Output: 8
 * Explanation: A -> B -> idle -> A -> B -> idle -> A -> B.
 *
 *
 *
 *
 * Note:
 *
 *
 * 	The number of tasks is in the range [1, 10000].
 * 	The integer n is in the range [0, 100].
 *
 *
 */

func TestTaskScheduler(t *testing.T) {
	var cases = []struct {
		input  []byte
		n      int
		output int
	}{
		{
			input:  []byte("AAABBB"),
			n:      2,
			output: 8,
		},
	}
	for _, c := range cases {
		x := leastInterval(c.input, c.n)
		if x != c.output {
			t.Fail()
		}
	}
}

// submission codes start here

func leastInterval(tasks []byte, n int) int {
	m := make([]int, 26)
	maxV := 0
	var ans int
	var maxCount int
	for _, v := range tasks {
		m[v-'A'] += 1
		if m[v-'A'] > m[maxV] {
			maxV = int(v - 'A')
		}
	}
	for _, v := range m {
		if v == m[maxV] {
			maxCount++
		}
	}
	ans = (m[maxV]-1)*(n+1) + maxCount
	if len(tasks) > ans {
		return len(tasks)
	}
	return ans
}

// submission codes end
