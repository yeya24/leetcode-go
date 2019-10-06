package tests

import (
	"testing"
)

/**
 * [547] Friend Circles
 *
 *
 * There are N students in a class. Some of them are friends, while some are not. Their friendship is transitive in nature. For example, if A is a direct friend of B, and B is a direct friend of C, then A is an indirect friend of C. And we defined a friend circle is a group of students who are direct or indirect friends.
 *
 *
 *
 * Given a N*N matrix M representing the friend relationship between students in the class. If M[i][j] = 1, then the ith and jth students are direct friends with each other, otherwise not. And you have to output the total number of friend circles among all the students.
 *
 *
 * Example 1:
 *
 * Input:
 * [[1,1,0],
 *  [1,1,0],
 *  [0,0,1]]
 * Output: 2
 * Explanation:The 0th and 1st students are direct friends, so they are in a friend circle. The 2nd student himself is in a friend circle. So return 2.
 *
 *
 *
 * Example 2:
 *
 * Input:
 * [[1,1,0],
 *  [1,1,1],
 *  [0,1,1]]
 * Output: 1
 * Explanation:The 0th and 1st students are direct friends, the 1st and 2nd students are direct friends, so the 0th and 2nd students are indirect friends. All of them are in the same friend circle, so return 1.
 *
 *
 *
 *
 * Note:
 *
 * N is in range [1,200].
 * M[i][i] = 1 for all students.
 * If M[i][j] = 1, then M[j][i] = 1.
 *
 *
 */

func TestFriendCircles(t *testing.T) {
	var cases = []struct {
		input  [][]int
		output int
	}{
		{
			input:  [][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}},
			output: 2,
		},
		{
			input:  [][]int{{1, 1, 0}, {1, 1, 1}, {0, 1, 1}},
			output: 1,
		},
		{
			input:  [][]int{{1, 0, 0, 1}, {0, 1, 1, 0}, {0, 1, 1, 1}, {1, 0, 1, 1}},
			output: 1,
		},
	}
	for _, c := range cases {
		x := findCircleNum2(c.input)
		if x != c.output {
			t.Fail()
		}
	}
}

// submission codes start here

func findCircleNum(M [][]int) int {
	n := len(M)
	c := 0
	visited := make([]int, n)
	for i := 0; i < n; i++ {
		if visited[i] == 0 {
			dfsFriends(M, visited, i)
			c++
		}
	}
	return c
}

func dfsFriends(M [][]int, visited []int, cur int) {
	n := len(M)
	if visited[cur] == 1 {
		return
	}
	visited[cur] = 1
	for i := 0; i < n; i++ {
		if M[cur][i] == 1 && visited[i] == 0 {
			dfsFriends(M, visited, i)
		}
	}
}

// Instead of using a array visited, use M[i][i] to check if is visited.
func findCircleNum2(M [][]int) int {
	n := len(M)
	c := 0
	for i := 0; i < n; i++ {
		if M[i][i] == 1 {
			dfsFriends2(M, i)
			c++
		}
	}
	return c
}

func dfsFriends2(M [][]int, cur int) {
	n := len(M)
	for i := 0; i < n; i++ {
		if M[cur][i] == 1 {
			M[cur][i] = 0
			M[i][cur] = 0
			dfsFriends2(M, i)
		}
	}
}

// submission codes end
