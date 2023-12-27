package main

import (
	"fmt"
	"strconv"
	"strings"
)

func romToAr(rawOne, rawTwo string) (int, int) {
	romanNum := map[string] int {
		"I":   1, "II":   2,
		"III": 3, "IV":   4,
		"V":   5, "VI":   6,
		"VII": 7, "VIII": 8,
		"IX":  9, "X":    10,
	}
	var numOne, numTwo int
	for key, value := range romanNum {
		if key == rawOne {numOne = value}
		if key == rawTwo {numTwo = value}
	}

	return numOne, numTwo
}

func arToRom (num int) string {
	var resalt strings.Builder 

	if num == 0 {
		resalt.WriteString("nulla")
	}

	for num > 0 {
		switch {
		case num >= 100:
			resalt.WriteString("C")
			num -= 100
		case num >= 90:
			resalt.WriteString("XC")
			num -= 90
		case num >= 50:
			resalt.WriteString("L")
			num -= 50
		case num >= 40:
			resalt.WriteString("XL")
			num -= 40
		case num >= 10:
			resalt.WriteString("X")
			num -= 10
		case num >= 5:
			resalt.WriteString("V")
			num -= 5
		case num >= 4:
			resalt.WriteString("IV")
			num -= 4	
		case num >= 1:
			resalt.WriteString("I")
			num -= 1
		
	}}
	return resalt.String()
}

func main (){
 
	var z, ok, rawOne, rawTwo string

	fmt.Println("Введите математическую операцию")
	for {
		count, _ := fmt.Scanf("%s %s %s %s", &rawOne, &z, &rawTwo, &ok)
		if count > 3 {
			fmt.Println("Неправильный формат математической операции.")
			return
		}else if count < 3 {
			fmt.Println("Недостаточно символов для математической операции или не используются пробелы между символами")
			return
		}
		
		numOne, err1 := strconv.Atoi(rawOne)
		numTwo, err2 := strconv.Atoi(rawTwo)

		flagR := false

		if (err1!= nil && err2 == nil) || (err2!= nil && err1 == nil ){
			fmt.Println("Используются разные системы счисления или неверные символы")
			return
		} else if err1!= nil && err2!= nil {
			flagR = true
			numOne, numTwo = romToAr(rawOne, rawTwo)
			if numOne == 0 || numTwo == 0 {
				fmt.Println("Используются неверные символы")
				return
			}
		}
		
		if numOne<=10 && numTwo<=10 && 0<numOne && 0<numTwo { 
			switch z {
			case "+":
				if flagR{
					fmt.Println(arToRom(numOne + numTwo))
				}else{fmt.Println(numOne + numTwo)}
			case "-":
				if flagR && numOne - numTwo < 0 {
					fmt.Println("В римской системе счисления нет отрицательных чисел")
				}else if flagR {fmt.Println(arToRom(numOne - numTwo))
				}else{fmt.Println(numOne - numTwo)}
			case "/":
				if flagR{
					fmt.Println(arToRom(numOne / numTwo))
				}else{fmt.Println(numOne / numTwo)}	
			case "*":
				if flagR{
					fmt.Println(arToRom(numOne * numTwo))
				}else{fmt.Println(numOne * numTwo)}
			default:
				fmt.Println("Присутствует недопустимый знак")
				return
		}} else {fmt.Println("Числа должны быть от 1 до 10")
			return}}
}