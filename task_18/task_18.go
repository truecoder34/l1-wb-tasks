package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
18.	Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
По завершению программа должна выводить итоговое значение счетчика.

1) Мьютекс
2) Атомики - https://gobyexample.com/atomic-counters
*/

func main() {
	// case 1
	cntr := InitializeCounterMutex()
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			cntr.IncrementMutexCounter()
		}()
	}
	wg.Wait()
	fmt.Printf("Counter after 10 goroutines : %v\n", cntr.counter.increment)
	// ------------------------------

	// case 2
	cntrAtomic := InitializeCounter()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			cntrAtomic.IncrementCounter()
		}()
	}
	wg.Wait()

	fmt.Printf("Counter Atomic after 10 goroutines : %v\n", cntrAtomic.increment)
	// ------------------------------

}

// case 1 . with Mutex
/*
	allows to call .Lock and Unlock functions to keep safe in goroutine
*/
type CounterMutex struct {
	sync.Mutex
	counter Counter
}

func InitializeCounterMutex() *CounterMutex {
	return &CounterMutex{} // constructor . return pointer on obj
}

func (c *CounterMutex) IncrementMutexCounter() {
	c.Lock()
	defer c.Unlock()
	c.counter.increment += 1
}

// ==================================================================
// case 2. atomics
type Counter struct {
	increment int32
}

func InitializeCounter() *Counter {
	return &Counter{}
}

/*
Reading atomics safely while they are being updated is also possible, using functions like atomic.LoadUint64.
*/
func (c *Counter) IncrementCounter() {
	atomic.AddInt32(&c.increment, 1) // add value to variable by address
}
