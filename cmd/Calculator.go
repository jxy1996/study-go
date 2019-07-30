package main

import (
	"fmt"
	Stack "github.com/jxy1996/study-go/dataStructure"
	"os"
	"strconv"
	"strings"
)

func main() {
	args := os.Args
	if args == nil && len(args) < 1 {
		fmt.Println("Null args..")
		return
	}
	ops := Stack.New()
	vals := Stack.New()
	str := strings.Split(args[1], "")
	for _, v := range str {
		switch v {
		case "(":
		case "+", "-", "*", "/":
			ops.Push(v)
		case ")":
			op := ops.Pop()
			v := vals.Pop().(int)
			switch op {
			case "+":
				v += vals.Pop().(int)
			case "-":
				v -= vals.Pop().(int)
			case "*":
				v *= vals.Pop().(int)
			case "/":
				v /= vals.Pop().(int)
			}
			vals.Push(v)
		default:
			s, err := strconv.Atoi(v)
			if err != nil {
				fmt.Printf("Atoi err:%v \n", err)
				return
			}
			vals.Push(s)
		}
	}
	//打印结果
	fmt.Println("Result:", vals.Peek())

}
