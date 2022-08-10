//З набору слів вам потрібно знайти слово, яке набрало найвищий бал.
//Кожна літера слова отримує бали відповідно до її позиції в алфавіті: a = 1, b = 2, c = 3 тощо.
//Вам потрібно повернути слово з найвищим балом у вигляді рядка.
//Якщо два слова мають однакові результати, повертається слово, яке з’являється першим у вихідному рядку.
//Усі літери будуть малими, і всі введені дані будуть дійсними.

package main

import (
	"fmt"
	"strings"
)

func main() {
	arr := [...]string{
		"a",
		"b",
		"c",
		"man i need a taxi up to ubud",
		"what time are we climbing up the volcano",
		"volcano",
		"take me to semynak",
		"aa b",
		"b aa",
		"bb d",
		"d bb",
		"aaa b",
	}
	for _, v := range arr {
		fmt.Println(High(v))
	}
}

func High(s string) string {
	wArr := strings.Split(s, " ")
	fin, max := 0, 0
	for in, val := range wArr {
		bArr := []byte(val)
		cur := 0
		for _, bval := range bArr {
			cur += int(bval) - 96
		}
		if cur > max {
			fin, max = in, cur
		}
		fmt.Println(cur)
	}
	return wArr[fin]
}
