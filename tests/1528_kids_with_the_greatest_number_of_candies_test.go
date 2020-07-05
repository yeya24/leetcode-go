package tests

import (
    "reflect"
    "testing"
)

/**
 * [1528] Kids With the Greatest Number of Candies
 *
 * Given the array candies and the integer extraCandies, where candies[i] represents the number of candies that the ith kid has.
 * 
 * For each kid check if there is a way to distribute extraCandies among the kids such that he or she can have the greatest number of candies among them. Notice that multiple kids can have the greatest number of candies.
 * 
 *  
 * Example 1:
 * 
 * 
 * Input: candies = [2,3,5,1,3], extraCandies = 3
 * Output: [true,true,true,false,true] 
 * Explanation: 
 * Kid 1 has 2 candies and if he or she receives all extra candies (3) will have 5 candies --- the greatest number of candies among the kids. 
 * Kid 2 has 3 candies and if he or she receives at least 2 extra candies will have the greatest number of candies among the kids. 
 * Kid 3 has 5 candies and this is already the greatest number of candies among the kids. 
 * Kid 4 has 1 candy and even if he or she receives all extra candies will only have 4 candies. 
 * Kid 5 has 3 candies and if he or she receives at least 2 extra candies will have the greatest number of candies among the kids. 
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: candies = [4,2,1,1,2], extraCandies = 1
 * Output: [true,false,false,false,false] 
 * Explanation: There is only 1 extra candy, therefore only kid 1 will have the greatest number of candies among the kids regardless of who takes the extra candy.
 * 
 * 
 * Example 3:
 * 
 * 
 * Input: candies = [12,1,12], extraCandies = 10
 * Output: [true,false,true]
 * 
 * 
 *  
 * Constraints:
 * 
 * 
 * 	2 <= candies.length <= 100
 * 	1 <= candies[i] <= 100
 * 	1 <= extraCandies <= 50
 * 
 */

func TestKidsWiththeGreatestNumberofCandies(t *testing.T) {
    var cases = []struct {
    	input []int
    	extra int
        output []bool
    }{
    	{
    	    input: []int{2,3,5,1,3},
    	    extra: 3,
    	    output: []bool{true, true, true, false, true},
        },
    }
    for _, c := range cases {
        x := kidsWithCandies(c.input, c.extra)
        if !reflect.DeepEqual(x, c.output) {
            t.Fail()
        }
    }
}

// submission codes start here

func kidsWithCandies(candies []int, extraCandies int) []bool {
	maxInit := 0
    for _, v := range candies {
        maxInit = max(maxInit, v)
    }

    res := make([]bool, len(candies))
    for i, v := range candies {
        if v + extraCandies >= maxInit {
            res[i] = true
        } else {
            res[i] = false
        }
    }
    return res
}

// submission codes end
