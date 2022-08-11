package main

import (
	"fmt"
	"time"
)

/*
25.
	Реализовать собственную функцию sleep.
*/

func main() {
	var sleepTime time.Duration
	/*
		Решение через Duration -
		A Duration represents the elapsed time between two instants as an int64 nanosecond count.

		временной отрезок
	*/
	sleepTime = time.Second * 10 // будем спать 10 секукнд
	fmt.Printf("We will sleep for %s \n", sleepTime)
	fmt.Println("Start now.")
	sleepNative(sleepTime)
	fmt.Println("Sleep finished.")

}

/**/
func sleepNative(duration time.Duration) {
	// count how when we need to wake up
	endTime := time.Duration(time.Now().UnixNano()) + duration
	for time.Duration(time.Now().UnixNano()) < endTime { // спим пока не настанет время проснуться
	}
}
