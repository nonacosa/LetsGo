package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func convertTpBin(n int) string {
	result := ""
	// 起始条件 结束条件  递增
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result

}

func printFile(filename string) {
	file, error := os.Open(filename)
	if error != nil {
		panic(error)
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	fmt.Println(
		convertTpBin(5),
		convertTpBin(13),
		convertTpBin(255),
		)
	printFile("abc.txt")
}
