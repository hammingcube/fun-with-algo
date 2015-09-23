package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

var Stdin *bufio.Reader

func init() {
	Stdin = bufio.NewReader(os.Stdin)
}

func main() {
	var s string

	for {
		count := map[rune]int{}
		isPalindrome := 0
		_, err := fmt.Fscanln(Stdin, &s)
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Encountered error:", err)
			break
		}
		for _, ch := range s {
			count[ch] += 1
		}
		for _, v := range count {
			//fmt.Printf("%s: %d\n", string(k), v)
			if v%2 != 0 {
				isPalindrome += 1
			}
		}
		if isPalindrome > 1 {
			fmt.Println(0)
		} else {
			fmt.Println(1)
		}
	}
}
