package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

/*
5.
	Разработать программу, которая будет последовательно отправлять значения в канал,
	а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.
*/

func main() {
	// 0 - read timing to work
	fmt.Println("Input lifetime of program")
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	text = strings.TrimSuffix(text, "\n") // remove symbols like \t \r
	text = strings.TrimSuffix(text, "\r") // remove symbols like \t \r
	lifetime, err := strconv.Atoi(text)
	if err != nil {
		panic(err)
	}

	// 1 - create msg channel
	messageChannel := make(chan string)
	defer close(messageChannel)

	// 2- start timer
	timeFin := time.Second * time.Duration(lifetime)
	timerStart := time.Now()

	// 3 - read data from channel
	go readData(messageChannel)

	// 4 - send data to channel while time left
	for time.Since(timerStart) < timeFin {
		messageChannel <- messageGenerate()

		fmt.Printf("Time left - %d \n", timeFin-time.Duration(time.Since(timerStart)))
		time.Sleep(time.Millisecond * 1000)

	}
}

/*
	Read all available data in channel
*/
func readData(messageCh chan string) {
	for data := range messageCh {
		fmt.Printf("Received message: %s\n", data)
	}
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
