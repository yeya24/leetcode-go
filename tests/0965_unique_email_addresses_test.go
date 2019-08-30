package tests

import (
	"testing"
)

/**
 * [965] Unique Email Addresses
 *
 * Every email consists of a local name and a domain name, separated by the @ sign.
 *
 * For example, in alice@leetcode.com, alice is the local name, and leetcode.com is the domain name.
 *
 * Besides lowercase letters, these emails may contain '.'s or '+'s.
 *
 * If you add periods ('.') between some characters in the local name part of an email address, mail sent there will be forwarded to the same address without dots in the local name.  For example, "alice.z@leetcode.com" and "alicez@leetcode.com" forward to the same email address.  (Note that this rule does not apply for domain names.)
 *
 * If you add a plus ('+') in the local name, everything after the first plus sign will be ignored. This allows certain emails to be filtered, for example m.y+name@email.com will be forwarded to my@email.com.  (Again, this rule does not apply for domain names.)
 *
 * It is possible to use both of these rules at the same time.
 *
 * Given a list of emails, we send one email to each address in the list.  How many different addresses actually receive mails?
 *
 *
 *
 *
 * Example 1:
 *
 *
 * Input: ["test.email+alex@leetcode.com","test.e.mail+bob.cathy@leetcode.com","testemail+david@lee.tcode.com"]
 * Output: 2
 * Explanation: "testemail@leetcode.com" and "testemail@lee.tcode.com" actually receive mails
 *
 *
 *
 *
 * Note:
 *
 *
 * 	1 <= emails[i].length <= 100
 * 	1 <= emails.length <= 100
 * 	Each emails[i] contains exactly one '@' character.
 * 	All local and domain names are non-empty.
 * 	Local names do not start with a '+' character.
 *
 *
 *
 */

func TestUniqueEmailAddresses(t *testing.T) {
	var cases = []struct {
		input  []string
		output int
	}{
		{
			input:  []string{"test.email+alex@leetcode.com", "test.e.mail+bob.cathy@leetcode.com", "testemail+david@lee.tcode.com"},
			output: 2,
		},
	}
	for _, c := range cases {
		x := numUniqueEmails(c.input)
		if x != c.output {
			t.Fail()
		}
	}
}

// submission codes start here

func numUniqueEmails(emails []string) int {
	m := make(map[string]struct{})
	c := 0
	for _, s := range emails {
		r := make([]byte, 0, len(s))
		i := 0
		var realEmail string
		for i < len(s) {
			if s[i] == '.' {
				i++
			} else if s[i] == '@' {
				realEmail = string(r) + s[i:]
				break
			} else if s[i] == '+' {
				for s[i] != '@' {
					i++
				}
			} else {
				r = append(r, s[i])
				i++
			}
		}
		if _, ok := m[realEmail]; !ok {
			m[realEmail] = struct{}{}
			c += 1
		}
	}
	return c
}

// submission codes end
