package main

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

func main() {
	var mapa sync.Map
	wg := sync.WaitGroup{} // init wait group
	// записываем 100 элементов
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			val := messageGenerate()
			fmt.Println(val)
			mapa.Store(strconv.Itoa(i), val)

		}(i)
	}

	time.Sleep(time.Millisecond * 1000)

	// читаем элементы
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			value, err := mapa.Load(strconv.Itoa(i))
			if err {
				log.Print(err)
			} else {
				//log.Printf("%s - %s", strconv.Itoa(i), value.(string))
				val, ok := value.(string)
				//fmt.Printf("%d - ",  ok, val)
				fmt.Println(val, ok)
				fmt.Println(ok)
				//log.Printf("%d - %s - ", i, val)
				// КАК ВЫВОДИТЬ ЧИСЛА А НЕ АДРЕС
			}
		}(i)
	}

	wg.Wait()
	//fmt.Println(mapa)
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

/*
	Generate random message from charset
*/
func messageGenerate() string {
	rand.Seed(time.Now().UnixNano())
	mesLen := rand.Int63n(int64(len(charset)))
	mes := make([]byte, mesLen) // place to store message
	// generate string symbol by symbol
	for i := range mes {
		time.Sleep(time.Millisecond * 10)
		mes[i] = charset[rand.Int63n(int64(len(charset)))]
	}
	return string(mes)
}
