package main

import (
	"fmt"
	"sync"
	"time"
)

/*
6.
	Реализовать все возможные способы остановки выполнения горутины.

	2 способа :
	закрыть каналом
	закрыть флагом
*/

func main() {

	// case 1 - сhannel close
	// quitChannel := make(chan bool)

	// go channelGoroutineClose(quitChannel)

	// time.Sleep(time.Second * 2)

	// quitChannel <- true
	// close(quitChannel)

	// case 2 - flag close
	quitFlag := false // flag allows to work

	wg := new(sync.WaitGroup) // CREATE WG FOR GOROUTINE
	wg.Add(1)                 // add 1 goroutine to workgroup

	go flagGoroutineClose(&quitFlag, wg)
	time.Sleep(time.Second * 2)
	quitFlag = true // flag allows to CLOSE

	wg.Wait()

}

/*
	Закрываем каналом
*/
func channelGoroutineClose(quitCh chan bool) {
	for {
		select {
		case <-quitCh:
			fmt.Println("Close goroutine via channel")
			time.Sleep(time.Millisecond * 100)
			return
		default:
			fmt.Println("goroutine continue to work")
			time.Sleep(time.Millisecond * 100)
		}
	}
}

/*
	Close via flag
*/
func flagGoroutineClose(quitFlag *bool, wg *sync.WaitGroup) {
	for {
		if *quitFlag {
			fmt.Println("Close goroutine via flag")
			time.Sleep(time.Millisecond * 1000)
			wg.Done()
			return
		} else {
			fmt.Println("goroutine continue to work")
			time.Sleep(time.Millisecond * 1000)
		}
	}
}
