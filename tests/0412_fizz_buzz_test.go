package tests

import (
	"reflect"
	"strconv"
	"testing"
)

/**
 * [412] Fizz Buzz
 *
 * Write a program that outputs the string representation of numbers from 1 to n.
 *
 * But for multiples of three it should output “Fizz” instead of the number and for the multiples of five output “Buzz”. For numbers which are multiples of both three and five output “FizzBuzz”.
 *
 * Example:
 *
 * n = 15,
 *
 * Return:
 * [
 *     "1",
 *     "2",
 *     "Fizz",
 *     "4",
 *     "Buzz",
 *     "Fizz",
 *     "7",
 *     "8",
 *     "Fizz",
 *     "Buzz",
 *     "11",
 *     "Fizz",
 *     "13",
 *     "14",
 *     "FizzBuzz"
 * ]
 *
 *
 */

func TestFizzBuzz(t *testing.T) {
	var cases = []struct {
		input  int
		output []string
	}{
		{
			input:  15,
			output: []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"},
		},
	}
	for _, c := range cases {
		x := fizzBuzz(c.input)
		if !reflect.DeepEqual(x, c.output) {
			t.Fail()
		}
	}
}

// submission codes start here

func fizzBuzz(n int) []string {
	s := make([]string, 0)
	for i := 1; i <= n; i++ {
		if i%15 == 0 {
			s = append(s, "FizzBuzz")
		} else if i%3 == 0 {
			s = append(s, "Fizz")
		} else if i%5 == 0 {
			s = append(s, "Buzz")
		} else {
			s = append(s, strconv.Itoa(i))
		}
	}
	return s
}

// submission codes end
