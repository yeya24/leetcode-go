package tests

import (
	"testing"
)

/**
 * [871] Keys and Rooms
 *
 * There are N rooms and you start in room 0.  Each room has a distinct number in 0, 1, 2, ..., N-1, and each room may have some keys to access the next room.
 *
 * Formally, each room i has a list of keys rooms[i], and each key rooms[i][j] is an integer in [0, 1, ..., N-1] where N = rooms.length.  A key rooms[i][j] = v opens the room with number v.
 *
 * Initially, all the rooms start locked (except for room 0).
 *
 * You can walk back and forth between rooms freely.
 *
 * Return true if and only if you can enter every room.
 *
 *
 *
 *
 * Example 1:
 *
 *
 * Input: [[1],[2],[3],[]]
 * Output: true
 * Explanation:
 * We start in room 0, and pick up key 1.
 * We then go to room 1, and pick up key 2.
 * We then go to room 2, and pick up key 3.
 * We then go to room 3.  Since we were able to go to every room, we return true.
 *
 *
 * Example 2:
 *
 *
 * Input: [[1,3],[3,0,1],[2],[0]]
 * Output: false
 * Explanation: We can't enter the room with number 2.
 *
 *
 * Note:
 *
 *
 * 	1 <= rooms.length <= 1000
 * 	0 <= rooms[i].length <= 1000
 * 	The number of keys in all rooms combined is at most 3000.
 *
 *
 */

func TestKeysandRooms(t *testing.T) {
	var cases = []struct {
		input  [][]int
		output bool
	}{
		{
			input:  [][]int{{1}, {2}, {3}, {}},
			output: true,
		},
		{
			input:  [][]int{{1, 3}, {3, 0, 1}, {2}, {0}},
			output: false,
		},
	}
	for _, c := range cases {
		x := canVisitAllRooms(c.input)
		if x != c.output {
			t.Fail()
		}
	}
}

// submission codes start here

func canVisitAllRooms(rooms [][]int) bool {
	visited := make([]bool, len(rooms))
	dfsRooms(rooms, visited, 0)
	for _, v := range visited {
		if v == false {
			return false
		}
	}
	return true
}

func dfsRooms(rooms [][]int, visited []bool, cur int) {
	if visited[cur] {
		return
	}
	visited[cur] = true
	for i := 0; i < len(rooms[cur]); i++ {
		if rooms[cur][i] != 0 {
			dfsRooms(rooms, visited, rooms[cur][i])
		}
	}
}

// submission codes end
