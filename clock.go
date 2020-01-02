package main

import (
	"fmt"
	"time"

	digs "github.com/niconosenzo/retroclock/vars"
)

var clock = [8][5]string{}

func main() {

	// main infinite loop

	for true {
		// clear the screen -> http://rosettacode.org/wiki/Terminal_control/Clear_the_screen#Go
		fmt.Print("\033[2J")

		now := time.Now()

		hours := now.Hour()
		minutes := now.Minute()
		seconds := now.Second()

		// for the hour (clock[0] and clock[1])

		clock[0] = digs.Digits[hours/10]
		clock[1] = digs.Digits[hours%10]

		// for the minutes (clock[3] and clock[4])

		clock[3] = digs.Digits[minutes/10]
		clock[4] = digs.Digits[minutes%10]

		// for the seconds (clock[6] and clock[7])

		clock[6] = digs.Digits[seconds/10]
		clock[7] = digs.Digits[seconds%10]

		// blink separators depending each 2 secs (depending if seconds are odd or given)
		if seconds%2 == 0 {
			// hh:mm separator
			clock[2] = digs.Separator[0]
			// mm:ss separator
			clock[5] = digs.Separator[0]
		} else {
			// hh:mm separator
			clock[2] = digs.Separator[1]
			// mm:ss separator
			clock[5] = digs.Separator[1]
		}

		// fmt.Println(hours, minutes, seconds)
		// fmt.Println(hours/10, minutes/10, seconds/10)
		// fmt.Println(hours%10, minutes%10, seconds%10)

		// for i := 0; i < 8; i++ {
		// 	for _, line := range clock[i] {
		// 		fmt.Print(line)
		// 	}
		// 	//fmt.Printf("%s", " ")
		// }

		for line := 0; line < len(clock[0]); line++ {
			for digit := range clock {
				fmt.Print(clock[digit][line], "  ")
			}
			fmt.Println()
		}

		time.Sleep(time.Second)

	}

}
