package tests

import (
    "reflect"
    "testing"
)

/**
 * [347] Top K Frequent Elements
 *
 * Given a non-empty array of integers, return the k most frequent elements.
 * 
 * Example 1:
 * 
 * 
 * Input: nums = [1,1,1,2,2,3], k = 2
 * Output: [1,2]
 * 
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: ums = [1], k = 1
 * Output: [1]
 * 
 * 
 * Note: 
 * 
 * 
 * 	You may assume k is always valid, 1 <= k <= number of unique elements.
 * 	Your algorithm's time complexity must be better than O(n log n), where n is the array's size.
 * 
 * 
 */

func TestTopKFrequentElements(t *testing.T) {
    var cases = []struct {
        input  []int
        k int
        output []int
    }{
        {
            input:  []int{1,1,1,2,2,3},
            k: 2,
            output: []int{1,2},
        },
        {
            input:  []int{1},
            k: 1,
            output: []int{1},
        },
    }
    for _, c := range cases {
        x := topKFrequent(c.input, c.k)
        if !reflect.DeepEqual(x, c.output) {
            t.Fail()
        }
    }
}

// submission codes start here

func topKFrequent(nums []int, k int) []int {
    m := make(map[int]int)
    maxFre := 0
    for _, v := range nums {
        m[v] +=1
        maxFre = max(maxFre, m[v])
    }
    bucket := make(map[int][]int, maxFre)
    for k, v := range m {
        bucket[v] = append(bucket[v], k)
    }
    res := []int{}
    for i := maxFre; i >= 1; i-- {
        p, ok := bucket[i]
        if !ok {
            continue
        }
        res = append(res, p...)
    }
    return res[:k]
}

// submission codes end
