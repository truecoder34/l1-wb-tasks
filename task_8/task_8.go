package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	value, err := inputInt64("Input int64 number to use : ")
	if err != nil {
		fmt.Printf("Failed to convert input to int64: %s \n", err)
	}

	bitIndex, err := inputInt64("Bit index to switch (1 is the first) ?")
	if err != nil {
		fmt.Printf("Failed to convert bit index input to int64: %s \n", err)
	}
	if bitIndex < 1 || bitIndex > 64 {
		fmt.Printf("Index should be from 1 to 64: %s \n", err)
		return
	}

	// case 1 solution
	fmt.Println("==========================================CASE 1==========================================")
	fmt.Printf("Value %d in bit format before switch: %064b \n", value, uint64(value))
	changedBitValue := switchBitStrSolution(value, bitIndex)
	if i, err := strconv.ParseInt(changedBitValue, 2, 64); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Value %d in bit format after switch: %s \n", i, changedBitValue)
	}
	fmt.Printf("Apple2Apple comparison : \n%064b \n%s \n", value, changedBitValue)

	// case 2 solution
	fmt.Println("==========================================CASE 2==========================================")
	fmt.Printf("Value %d in bit format before switch: %064b \n", value, uint64(value))
	changedBitValue2 := switchBit(value, bitIndex)
	fmt.Printf("Value %d in bit format after switch: %s \n", changedBitValue2, changedBitValue)
	fmt.Printf("Apple2Apple comparison : \n%064b \n%064b \n", value, changedBitValue2)
}

/*
	SOLUTION 1 - to string convert
*/
func switchBitStrSolution(val int64, position int64) string {
	// Only need it to detect bit Index to change !
	toBitStr := fmt.Sprintf("%064b", val)
	fmt.Printf("[Convert bit number to string format]: %s \n", toBitStr)

	bitsSlice := strings.Split(toBitStr, "")
	fmt.Printf("[Convert string to slice]: %v \n", bitsSlice)

	// switch bit by index
	if bitsSlice[int64(len(bitsSlice))-position] == "0" {
		bitsSlice[int64(len(bitsSlice))-position] = "1"
	} else {
		bitsSlice[int64(len(bitsSlice))-position] = "0"
	}
	// join slice back to string
	res := strings.Join(bitsSlice, "")
	fmt.Printf("[After joining back to string] %v\n", res)

	return res
}

/*
ParseInt always returns an int64 value.
Depending on bitSize, this value will fit into int, int8, int16, int32, or int64.
If the value cannot be represented by a signed integer of the size given by bitSize,
 then err.Err = ErrRange.
 https://stackoverflow.com/questions/50815512/when-casting-an-int64-to-uint64-is-the-sign-retained
*/
func switchBit(val int64, position int64) int64 {
	toBit := uint64(val)
	// need to check what bit 0 or 1 is onr position
	toBitStr := fmt.Sprintf("%064b", val)
	fmt.Printf("[Convert bit number to string format]: %s \n", toBitStr)
	bitsSlice := strings.Split(toBitStr, "")
	fmt.Printf("[Convert string to slice]: %v \n", bitsSlice)

	mask := uint64(1 << position) // make left shift to make 1 on index position
	fmt.Printf("[MASK FOR CHOSEN POSITION CONVERT]: %b \n", mask)

	// Make bit operations in DECIMAL representation
	if bitsSlice[int64(len(toBitStr))-position-1] == "0" {
		// val := toBit &^ mask
		// fmt.Sprintf("%064b", val)
		//return int64(val)
		return int64(toBit | mask)
	} else {
		return int64(toBit &^ mask)
	}
}

func inputInt64(line string) (int64, error) {
	fmt.Println(line)
	reader := bufio.NewReader(os.Stdin)
	number, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	number = strings.TrimSuffix(number, "\n") // remove symbols like \t \r
	number = strings.TrimSuffix(number, "\r") // remove symbols like \t \r

	n, err := strconv.ParseInt(number, 10, 64)
	if err != nil {
		fmt.Printf("Error on convert: %d of type %T \n", n, n)

	}

	return n, nil
}

// 0000000000000000000000000000000000000000000001010100010100000010
// 0000000000000000000000000000000000000000000001010100010100100010
