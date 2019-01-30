// Package twelve outputs the lyrics to The Twelve Days of Christmas
package twelve

import (
	"fmt"
)

var data = [][]string{
	[]string{"first", "a Partridge in a Pear Tree"},
	[]string{"second", "two Turtle Doves"},
	[]string{"third", "three French Hens"},
	[]string{"fourth", "four Calling Birds"},
	[]string{"fifth", "five Gold Rings"},
	[]string{"sixth", "six Geese-a-Laying"},
	[]string{"seventh", "seven Swans-a-Swimming"},
	[]string{"eighth", "eight Maids-a-Milking"},
	[]string{"ninth", "nine Ladies Dancing"},
	[]string{"tenth", "ten Lords-a-Leaping"},
	[]string{"eleventh", "eleven Pipers Piping"},
	[]string{"twelfth", "twelve Drummers Drumming"},
}

// Song outputs the lyrics to The Twelve Days of Christmas
func Song() string {
	song := ""
	for day := 1; day <= 12; day++ {
		song += Verse(day) + "\n"
	}
	return song
}

// Verse outputs the lyrics to a given verse of The Twelve Days of Christmas
func Verse(day int) string {
	i := day - 1
	quantity := data[i][0]
	items := list(day)

	return fmt.Sprintf("On the %s day of Christmas my true love gave to me: %s.", quantity, items)
}

func list(day int) string {
	if day == 1 {
		return data[0][1]
	}
	l := ""
	for d := day; d > 0; d-- {
		l += data[d-1][1]
		if d == 2 {
			l += ", and "
		} else if d > 2 {
			l += ", "
		}
	}
	return l
}
