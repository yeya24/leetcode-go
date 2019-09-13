package tests

import (
    "testing"
)

/**
 * [1289] Day of the Week
 *
 * Given a date, return the corresponding day of the week for that date.
 * 
 * The input is given as three integers representing the day, month and year respectively.
 * 
 * Return the answer as one of the following values {"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}.
 * 
 *  
 * Example 1:
 * 
 * 
 * Input: day = 31, month = 8, year = 2019
 * Output: "Saturday"
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: day = 18, month = 7, year = 1999
 * Output: "Sunday"
 * 
 * 
 * Example 3:
 * 
 * 
 * Input: day = 15, month = 8, year = 1993
 * Output: "Sunday"
 * 
 * 
 *  
 * Constraints:
 * 
 * 
 * 	The given dates are valid dates between the years 1971 and 2100.
 * 
 */

func TestDayoftheWeek(t *testing.T) {
    var cases = []struct {
    	day int
    	month int
    	year int
        output string
    }{
        {
        	day: 31,
        	month: 8,
        	year: 2019,
            output: "Saturday",
        },
        {
            day: 18,
            month: 7,
            year: 1999,
            output: "Sunday",
        },
        {
            day: 15,
            month: 8,
            year: 1993,
            output: "Sunday",
        },
    }
    for _, c := range cases {
        x := dayOfTheWeek(c.day, c.month, c.year)
        if x != c.output {
            t.Fail()
        }
    }
}

// submission codes start here

func dayOfTheWeek(day int, month int, year int) string {
    t := []int{0, 3, 2, 5, 0, 3, 5, 1, 4, 6, 2, 4}
    if month < 3 {
        year -= 1
    }
    w := (year + year/4 - year/100 + year/400 + t[month-1] + day) % 7
    switch w {
    case 1:
        return "Monday"
    case 2:
        return "Tuesday"
    case 3:
        return "Wednesday"
    case 4:
        return "Thursday"
    case 5:
        return "Friday"
    case 6:
        return "Saturday"
    case 0:
        return "Sunday"
    }
    return ""
}

// submission codes end
