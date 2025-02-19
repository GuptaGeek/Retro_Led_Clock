package main

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
)

type placeholder [5]string

func main() {

	
	for j := 0; j < 10; j++ {
		screen.Clear()
		screen.MoveTopLeft()
		RetroLedClock()
		time.Sleep(time.Second)
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
	now := time.Now()
	hour, minute, second := now.Hour(), now.Minute(), now.Second()
	clock := [8]int{hour / 10, hour % 10, 10, minute / 10, minute % 10, 10, second / 10, second % 10}
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
				if digit == 10 {
					fmt.Print(seperator[i])
				} else {
					fmt.Print(digits[digit][i], " ")
				}
			}
			fmt.Println()
		}
	}

}
