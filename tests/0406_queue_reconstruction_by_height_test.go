package tests

import (
    "fmt"
    "sort"
    "testing"
)

/**
 * [406] Queue Reconstruction by Height
 *
 * Suppose you have a random list of people standing in a queue. Each person is described by a pair of integers (h, k), where h is the height of the person and k is the number of people in front of this person who have a height greater than or equal to h. Write an algorithm to reconstruct the queue.
 * 
 * Note:
 * The number of people is less than 1,100.
 *  
 * 
 * Example
 * 
 * 
 * Input:
 * [[7,0], [4,4], [7,1], [5,0], [6,1], [5,2]]
 * 
 * Output:
 * [[5,0], [7,0], [5,2], [6,1], [4,4], [7,1]]
 * 
 * 
 *  
 * 
 */

func TestQueueReconstructionbyHeight(t *testing.T) {
    var cases = []struct {
        input  [][]int
        output [][]int
    }{
        {
            input:  [][]int{{7,0}, {4,4}, {7,1}, {5,0}, {6,1}, {5,2}},
            output: [][]int{{5,0}, {7,0}, {5,2}, {6,1}, {4,4}, {7,1}},
        },
    }
    for _, c := range cases {
        x := reconstructQueue(c.input)
        fmt.Println(x)
    }
}

// submission codes start here

func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func (i, j int) bool {
        if people[i][0] != people[j][0] {
            return people[i][0] > people[j][0]
        }
        return people[i][1] < people[j][1]
    })
	res := make([][]int, len(people))
	for i := 0; i < len(people); i++ {
	    p := people[i][1]
	    copy(res[p+1:], res[p:])
	    res[p] = people[i]
    }
   return res
}

// submission codes end
