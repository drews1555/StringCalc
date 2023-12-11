package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var input []string
	var operator string
	fmt.Println("Введите выражение")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	switch {
	case strings.Contains(scanner.Text(), " + ") == true:
		input = strings.Split(scanner.Text(), " + ")
		operator = "+"
		osh1(input[1])
	case strings.Contains(scanner.Text(), " - ") == true:
		input = strings.Split(scanner.Text(), " - ")
		operator = "-"
		osh1(input[1])
	case strings.Contains(scanner.Text(), " * ") == true:
		input = strings.Split(scanner.Text(), " * ")
		operator = "*"
		osh2(input[1])
	case strings.Contains(scanner.Text(), " / ") == true:
		input = strings.Split(scanner.Text(), " / ")
		operator = "/"
		osh2(input[1])
	default:
		fmt.Println("Неверная арифметическая операция")
		os.Exit(1)
	}
	a := input[0]
	b := input[1]
	//fmt.Println(a, b)

	if _, err := strconv.Atoi(a); err == nil {
		fmt.Println("ВВедите строку первым аргументом")
		os.Exit(1)
	}
	switch operator {
	case "+":
		c := a[1 : len(a)-1]
		d := b[1 : len(b)-1]
		osh10(c)
		osh10(d)
		result := c + d
		fmt.Println("\"" + result + "\"")
	case "-":
		c := a[1 : len(a)-1]
		d := b[1 : len(b)-1]
		osh10(c)
		osh10(d)
		e := strings.Contains(c, d)
		if e == true {
			g := strings.Split(c, d)
			result := g[0] + g[1]
			//result := c[:len(c)-(len(d))]
			fmt.Println("\"" + result + "\"")
		} else {
			fmt.Println("\"" + c + "\"")
		}
	case "*":
		c := a[1 : len(a)-1]
		osh10(c)
		var d string
		if c, err := strconv.Atoi(b); err == nil {
			osh110(c)
			for i := 0; i < c; i++ {
				d += a[1 : len(a)-1]
			}
			if len(d) > 40 {
				fmt.Println("\"" + d[0:40] + "..." + "\"")
			} else {
				fmt.Println("\"" + d + "\"")
			}
		} else {
			fmt.Print("Операции умножения и деления проводятся только с целыми числами")
		}
	case "/":
		c := a[1 : len(a)-1]
		osh10(c)
		if c, err := strconv.Atoi(b); err == nil {
			osh110(c)
			fmt.Println("\"" + a[1:((len(a)-2)/c)+1] + "\"")
		} else {
			fmt.Println("Операции умножения и деления проводятся только с целыми числами")
		}
	}
}

func osh10(a string) {
	//Строка не длинее 10
	if len(a) > 10 {
		fmt.Println("Длинна строк не должна превышать 10 символов")
		os.Exit(1)
	}
}

func osh110(a int) {
	if a > 10 || a < 1 {
		fmt.Println("Введите числа от 1 до 10")
		os.Exit(1)
	}
}
func osh1(a string) {
	if strings.Contains(a, "\"") {
	} else {
		fmt.Println("Неподходящая операция")
		os.Exit(1)
	}
}

func osh2(a string) {
	if strings.Contains(a, "\"") {
		fmt.Println("Неподходящая операция")
		os.Exit(1)
	}
}
