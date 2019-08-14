package tests

import (
	"strings"
	"testing"
)

/**
 * [1205] Defanging an IP Address
 *
 * Given a valid (IPv4) IP address, return a defanged version of that IP address.
 *
 * A defanged IP address replaces every period "." with "[.]".
 *
 *
 * Example 1:
 * Input: address = "1.1.1.1"
 * Output: "1[.]1[.]1[.]1"
 * Example 2:
 * Input: address = "255.100.50.0"
 * Output: "255[.]100[.]50[.]0"
 *
 *
 * Constraints:
 *
 *
 * 	The given address is a valid IPv4 address.
 *
 */

func TestDefanginganIPAddress(t *testing.T) {
	var cases = []struct {
		input  string
		output string
	}{
		{
			input:  "1.1.1.1",
			output: "1[.]1[.]1[.]1",
		},
		{
			input:  "255.100.50.0",
			output: "255[.]100[.]50[.]0",
		},
	}
	for _, c := range cases {
		x := defangIPaddr(c.input)
		if x != c.output {
			t.Fail()
		}
	}
}

// submission codes start here

func defangIPaddr(address string) string {
	return strings.Replace(address, ".", "[.]", -1)
}

// submission codes end
