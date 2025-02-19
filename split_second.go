package main

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
)

type placeholder [5]string
const splitsecond=time.Second/10

func main() {

	for j := 0; j < 10; j++ {
		screen.Clear()
		screen.MoveTopLeft()
		RetroLedClock()
		time.Sleep(splitsecond)
	}

}
func RetroLedClock() {

	digits := [10]placeholder{

		[5]string{"███", "█ █", "█ █", "█ █", "███"},
		[5]string{"  █", "  █", "  █", "  █", "  █"},
		[5]string{"███", "  █", "███", "█  ", "███"},
		[5]string{"███", "  █", "███", "  █", "███"},
		[5]string{"█ █", "█ █", "███", "  █", "  █"},
		[5]string{"███", "█  ", "███", "  █", "███"},
		[5]string{"███", "█  ", "███", "█ █", "███"},
		[5]string{"███", "  █", "  █", "  █", "  █"},
		[5]string{"███", "█ █", "███", "█ █", "███"},
		[5]string{"███", "█ █", "███", "  █", "███"},
	}
	alarm := [8]placeholder{
		[5]string{" ", " ", " ", " ", " "},
		[5]string{"███", "█ █", "███", "█ █", "█ █"},
		[5]string{"█  ", "█  ", "█  ", "█  ", "███"},
		[5]string{"███", "█ █", "███", "█ █", "█ █"},
		[5]string{"███", "█ █", "██ ", "█ █", "█ █"},
		[5]string{"█ █", "███", "█ █", "█ █", "█ █"},
		[5]string{"█  ", "█  ", "█  ", " ", "█  "},
		[5]string{" ", " ", " ", " ", " "},
	}
	seperator := placeholder{"   ", " █ ", "   ", " █ ", "   "}
	spaces := placeholder{"   ", "   ", "   ", "   ", "   "}
	dot := placeholder{"   ", "   ", "   ", "   ", " █ "}
	now := time.Now()
	hour, minute, second := now.Hour(), now.Minute(), now.Second()
	ssec := now.Nanosecond() / 100000000
	clock := [...]int{hour / 10, hour % 10, 10, minute / 10, minute % 10, 10, second / 10, second % 10, 11, ssec}
	if second%10 == 0 {
		for i := 0; i < 5; i++ {
			for j := 0; j < 8; j++ {
				fmt.Print(alarm[j][i], " ")
			}
			fmt.Println()
		}
	} else {

		for i := 0; i <= 4; i++ {
			for _, digit := range clock {
				if digit == 10 && second%2 != 0 {
					fmt.Print(seperator[i])
				} else if digit == 10 && second%2 == 0 {
					fmt.Print(spaces[i])
				} else if digit == 11 && second%2 != 0 {
					fmt.Print(dot[i])
				} else if digit == 11 && second%2 == 0 {
					fmt.Print(spaces[i])
				} else {
					fmt.Print(digits[digit][i], " ")
				}
			}
			fmt.Println()
		}
	}

}
