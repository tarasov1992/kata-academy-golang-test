package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Expression struct {
	X, Y     int
	Operator string
}

func (exp Expression) Add() int {
	return exp.X + exp.Y
}

func (exp Expression) Subtract() int {
	return exp.X - exp.Y
}

func (exp Expression) Multiply() int {
	return exp.X * exp.Y
}

func (exp Expression) Divide() int {
	return exp.X / exp.Y
}

func main() {
	// declaring a var that will read and store input from stdin
	buffer := bufio.NewReader(os.Stdin)

	welcomeMessage := "What do you want to calculate?"

	for {
		fmt.Println(welcomeMessage)
		usrInputStr, _ := buffer.ReadString('\n')
		usrInputStr = strings.TrimSpace(usrInputStr)
		usrInputNum, _ := strconv.Atoi(usrInputStr)
		fmt.Println(usrInputNum + 8)
	}
}
