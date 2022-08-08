package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

/*
	Реализовать конкурентную запись данных в map.
*/

/*
	Структура со своим собсвтенным мьютексом
*/
type SafeMap struct {
	mutex         sync.RWMutex
	randomDataMap map[string]string
}

/*
	методы для конкерутной записи и чтения данных в мапу
	Сеттер
	методы Get() и Store() должны быть определены для указателя на Counters,
	а не на Counters (тоесть не func (c Counters) Load(key string) int { ... },
	потому что в таком случае значение ресивера (c) копируется,
	вместе с чем скопируется и мьютекс в нашей структуре,
	что лишает всю затею смысла и приводит к проблемам.
*/
func (sm *SafeMap) AddValue(textKey string, textVal string) {
	sm.mutex.Lock()
	defer sm.mutex.Unlock()
	sm.randomDataMap[textKey] = textVal
}

// Геттер
func (sm *SafeMap) GetValue(textKey string) (string, error) {
	sm.mutex.RLock() // Если метод нуждается только в чтении — он использует RLock(),
	//который не заблокирует другие операции чтения, но заблокирует операцию записи и наоборот
	defer sm.mutex.RUnlock()
	if elem, ok := sm.randomDataMap[textKey]; ok {
		return elem, nil
	}
	return "0", errors.New("Value does not exists")
}

func ConstructorSafeMap() *SafeMap {
	return &SafeMap{
		randomDataMap: make(map[string]string),
	}
}

func main() {
	wg := sync.WaitGroup{}       // init wait group
	mapa := ConstructorSafeMap() // init structure of safe map
	// записываем 100 элементов
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			mapa.AddValue(strconv.Itoa(i), messageGenerate())
		}(i)
	}

	time.Sleep(time.Millisecond * 1000)

	// читаем элементы
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			value, err := mapa.GetValue(strconv.Itoa(i))
			if err != nil {
				log.Print(err)
			} else {
				log.Printf("%s - %s", strconv.Itoa(i), value)
			}
		}(i)
	}

	wg.Wait()
	fmt.Println(mapa)
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
