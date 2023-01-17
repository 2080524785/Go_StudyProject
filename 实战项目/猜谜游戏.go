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

func v1() {
	maxNum := 100
	secretNumber := rand.Intn(maxNum)
	fmt.Println(secretNumber)
}
func v2() {
	maxNum := 100
	// 设计随机数种子
	rand.Seed(time.Now().UnixNano())
	secretNumber := rand.Intn(maxNum)
	fmt.Println(secretNumber)
}
func v3() {
	maxNum := 100
	// 设计随机数种子
	rand.Seed(time.Now().UnixNano())
	secretNumber := rand.Intn(maxNum)
	fmt.Println(secretNumber)
}
func v4() {
	maxNum := 100
	// 设计随机数种子
	rand.Seed(time.Now().UnixNano())
	secretNumber := rand.Intn(maxNum)
	fmt.Println(secretNumber)
	fmt.Println("please input your guess")

	reader := bufio.NewReader(os.Stdin)
	for {

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("please try again", err)
			continue
		}
		// 返回没有提供的尾随后缀字符串的 s。如果 s 不以 suffix 结尾，则 s 原样返回。
		input = strings.TrimSuffix(input, "\n")

		guess, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("please try again")
			continue
		}
		fmt.Println("Your guess is", guess)
		if guess > secretNumber {
			fmt.Println(">")
		} else if guess < secretNumber {
			fmt.Println("<")
		} else {
			fmt.Println("=")
			break
		}
	}
}

// 更换输入方式
func v5() {
	maxNum := 100
	// 设计随机数种子
	rand.Seed(time.Now().UnixNano())
	secretNumber := rand.Intn(maxNum)
	fmt.Println(secretNumber)
	fmt.Println("please input your guess")

	reader := bufio.NewReader(os.Stdin)
	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("please try again", err)
			continue
		}
		// 返回没有提供的尾随后缀字符串的 s。如果 s 不以 suffix 结尾，则 s 原样返回。
		input = strings.TrimSuffix(input, "\n")

		guess, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("please try again")
			continue
		}
		fmt.Println("Your guess is", guess)
		if guess > secretNumber {
			fmt.Println(">")
		} else if guess < secretNumber {
			fmt.Println("<")
		} else {
			fmt.Println("=")
			break
		}
	}
}
func main() {
	v5()
}
