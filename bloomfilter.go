package main

import (
	"fmt"
	"hash/fnv"
	"math"
)

func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

func bloomInit(itemNum uint32, falsePositive float64) (uint32, uint32) {
	var bitArrSize float64
	var numHashFunc float64

	bitArrSize = -(float64(itemNum) * math.Log(falsePositive)) / math.Pow(math.Log(2), 2)
	numHashFunc = (bitArrSize / float64(itemNum)) * math.Log(float64(2))

	return uint32(math.Ceil(bitArrSize)), uint32(math.Ceil(numHashFunc))
}

func bloomAdd(bitArr []bool, bitArrSize uint32, numHashFunc uint32, word string) {
	var i uint32
	for i = 0; i < numHashFunc; i++ {
		bitArr[hash(word+string(i))%bitArrSize] = true
	}
}

func bloomCheck(bitArr []bool, bitArrSize uint32, numHashFunc uint32, word string) bool {
	var i uint32
	for i = 0; i < numHashFunc; i++ {
		if bitArr[hash(word+string(i))%bitArrSize] != true {
			return false
		}
	}
	return true
}

// func bloomFree() {
// 	//delete
// }

func main() {

	itemNum := 100000
	falsePositive := 0.001
	bitArrSize, numHashFunc := bloomInit(uint32(itemNum), falsePositive)

	bitArr := make([]bool, bitArrSize) // in Go all values are false by default

	var myWord, myOp string
	for true {
		fmt.Println("Enter ADD for adding a key, CHECK for searching a key or QUIT:")
		fmt.Scanln(&myOp)

		if myOp == "ADD" {
			fmt.Println("Enter a key to add:")
			fmt.Scanln(&myWord)

			bloomAdd(bitArr, bitArrSize, numHashFunc, myWord)

		} else if myOp == "CHECK" {
			fmt.Println("Enter a key to check:")
			fmt.Scanln(&myWord)

			if bloomCheck(bitArr, bitArrSize, numHashFunc, myWord) == true {
				fmt.Printf("%s exists.\n", myWord)
			} else {
				fmt.Printf("%s does not exist.\n", myWord)
			}
		} else if myOp == "QUIT" {
			break
		} else {
			fmt.Println("ADD or CHECK or QUIT?:")
		}
	}
}
