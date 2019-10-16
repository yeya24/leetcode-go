package tests

import (
	"math"
	"testing"
)

/**
 * [8] String to Integer (atoi)
 *
 * Implement atoi which converts a string to an integer.
 *
 * The function first discards as many whitespace characters as necessary until the first non-whitespace character is found. Then, starting from this character, takes an optional initial plus or minus sign followed by as many numerical digits as possible, and interprets them as a numerical value.
 *
 * The string can contain additional characters after those that form the integral number, which are ignored and have no effect on the behavior of this function.
 *
 * If the first sequence of non-whitespace characters in str is not a valid integral number, or if no such sequence exists because either str is empty or it contains only whitespace characters, no conversion is performed.
 *
 * If no valid conversion could be performed, a zero value is returned.
 *
 * Note:
 *
 *
 * 	Only the space character ' ' is considered as whitespace character.
 * 	Assume we are dealing with an environment which could only store integers within the 32-bit signed integer range: [-2^31,  2^31 - 1]. If the numerical value is out of the range of representable values, INT_MAX (2^31 - 1) or INT_MIN (-2^31) is returned.
 *
 *
 * Example 1:
 *
 *
 * Input: "42"
 * Output: 42
 *
 *
 * Example 2:
 *
 *
 * Input: "   -42"
 * Output: -42
 * Explanation: The first non-whitespace character is '-', which is the minus sign.
 *              Then take as many numerical digits as possible, which gets 42.
 *
 *
 * Example 3:
 *
 *
 * Input: "4193 with words"
 * Output: 4193
 * Explanation: Conversion stops at digit '3' as the next character is not a numerical digit.
 *
 *
 * Example 4:
 *
 *
 * Input: "words and 987"
 * Output: 0
 * Explanation: The first non-whitespace character is 'w', which is not a numerical
 *              digit or a +/- sign. Therefore no valid conversion could be performed.
 *
 * Example 5:
 *
 *
 * Input: "-91283472332"
 * Output: -2147483648
 * Explanation: The number "-91283472332" is out of the range of a 32-bit signed integer.
 *              Thefore INT_MIN (-2^31) is returned.
 *
 */

func TestStringtoInteger(t *testing.T) {
	var cases = []struct {
		input  string
		output int
	}{
		{
			input:  "42",
			output: 42,
		},
		{
			input:  "   -42",
			output: -42,
		},
		{
			input:  "4193 with words",
			output: 4193,
		},
		{
			input:  "words and 987",
			output: 0,
		},
		{
			input:  "-91283472332",
			output: -2147483648,
		},
	}
	for _, c := range cases {
		x := myAtoi(c.input)
		if x != c.output {
			t.Fail()
		}
	}
}

// submission codes start here

func myAtoi(str string) int {
	res := 0
	isMinus := false
	isOver := false
	start := false
	for _, v := range str {
		if (v != ' ' && !start) || start {
			switch checkByte(byte(v)) {
			case 0:
				start = true
				res = res*10 + int(v-'0')
				if res > math.MaxInt32 {
					isOver = true
					goto END
				}
			case 1:
				if start {
					goto END
				} else {
					if v == '-' {
						isMinus = true
					}
					start = true
				}
			default:
				goto END
			}
		}
	}
END:
	if isMinus {
		if isOver {
			return math.MinInt32
		}
		return -res
	}
	if isOver {
		return math.MaxInt32
	}
	return res
}

func checkByte(b byte) int {
	if b >= '0' && b <= '9' {
		return 0
	}
	if b == '-' || b == '+' {
		return 1
	}
	return -1
}

// submission codes end
