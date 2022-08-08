package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

/*
Реализовать постоянную запись данных в канал (главный поток).
Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout. Необходима возможность выбора количества воркеров при старте.

Программа должна завершаться по нажатию Ctrl+C.
Выбрать и обосновать способ завершения работы всех воркеров.
*/

func main() {
	fmt.Printf("Enter number of workers: ")
	// var reader = bufio.NewReader(os.Stdin)
	// n, _ := reader.ReadString('\n')
	// N, err := strconv.Atoi(n)
	// if err != nil {
	// 	panic(err)
	// }
	//fmt.Printf("N = %d ", N)
	var N int
	_, err := fmt.Scanf("%d", &N)
	if err != nil {
		fmt.Println("Not number received ! ", err)
	}

	// 1 - контроль за прерываением
	signalChannel := make(chan os.Signal, 1) // channel to control signals from keyboard
	signal.Notify(signalChannel, syscall.SIGINT)
	go stopControl(signalChannel)

	// 2 - запуск записи в канал
	exec(N)
}

/*
	Break control in function execution
	Stop symbol : Ctrl + C

	Ctrl-C is SIGINT.
	Ctrl-\ is SIGQUIT (by default)
*/
func stopControl(signalCh chan os.Signal) {
	for {
		sig := <-signalCh // get signal from channel. Waiting for Crtl+C
		switch sig {
		case syscall.SIGINT: // ctrl + c case handling
			fmt.Printf("Break request received. \n")
			os.Exit(0)
		default: // all other cases - go next
			fmt.Printf("Unknown command/signal received. \n")
		}
	}
}

/*
	execution function to process data in channels
*/
func exec(workersNumber int) {
	// 0 - создаем горутину генерирующую  сообщения
	// make channel for data processing
	dataChannel := make(chan string, workersNumber)
	go seedData2Channel(dataChannel)

	// 1 - создаем воркеров для чтения из канала и записи в поток
	var wg sync.WaitGroup // создаем группу горутин. один воркер - одна группа
	for i := 0; i < workersNumber; i++ {
		workerIndex := i
		wg.Add(1) // группа будет состоять из 1ой горутины. определяет значение внутреннего счетчика активных элементов
		go func() {
			defer wg.Done() // сигнализировать, что элемент группы завершил свое выполнение. в горутине вызываем DONE. уменьшает внутренний счетчик активных элементов на единицу.
			for data := range dataChannel {
				fmt.Printf("[WORKER-%d] DATA: %s \n", workerIndex, data)
				time.Sleep(time.Millisecond * 1000)
			}
		}()
	}
	wg.Wait() // ожидаем завершение всех горутин.
	//Метод деблокирует функцию main, когда внутренний счетчик активных элементов в группе wg стает равен 0
}

/*
	Push data to channel
*/
func seedData2Channel(dataCh chan string) {
	// write 10k messages to channel
	for i := 0; i < 10000; i++ {
		dataCh <- messageGenerate()
		time.Sleep(time.Millisecond * 1000)
	}
	close(dataCh)
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

/*
	Generate random message from charset
*/
func messageGenerate() string {
	rand.Seed(time.Now().UnixNano())
	mesLen := rand.Intn(len(charset)) // random string length
	mes := make([]byte, mesLen)       // place to store message
	// generate string symbol by symbol
	for i := range mes {
		mes[i] = charset[rand.Intn(len(charset))]
	}
	return string(mes)
}
