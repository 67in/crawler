package main

import (
	"regexp"
	"fmt"
)

func main() {
	//b1 := regexp.MustCompile(`123.`).MatchString(`123*embda`)
	//fmt.Println(b1)
	//
	//b2 := regexp.MustCompile(`123....`).MatchString(`123memeda`)
	//fmt.Println(b2)
	//
	//b3 := regexp.MustCompile(`[abc][dc]`).MatchString(`saca`)
	//fmt.Println(b3)
	//
	//b4 := regexp.MustCompile(`a[abc]`).MatchString(`ac`)
	//fmt.Println(b4)
	//
	//b5 := regexp.MustCompile(`[a-x]`).MatchString(`aC`)
	//fmt.Println(b5)
	//
	//b6 := regexp.MustCompile(`[a-zA-Z][a-z]c`).MatchString(`aac`)
	//fmt.Println(b6)
	//
	//b7 := regexp.MustCompile(`[1-9]`).MatchString(`120`)
	//fmt.Println(b7)
	//
	//b8 := regexp.MustCompile(`[^abc]`).MatchString(`ADBIDJAID%$%$ABabc`)
	//fmt.Println(b8)
	//
	//b9 := regexp.MustCompile(`[^a-z]`).MatchString(`Aamemeda`) //
	//fmt.Println(b9)
	//
	//b10 := regexp.MustCompile(`1[^2345]`).MatchString(`1a6`) //
	//fmt.Println(b10)
	//
	//b11 := regexp.MustCompile(`^\d`).MatchString(`123`)
	//fmt.Println(b11)
	//
	//b12 := regexp.MustCompile(`\d\d`).MatchString(`12bc`) //[0-9]
	//fmt.Println(b12)
	//
	//b13 := regexp.MustCompile(`\D`).MatchString(`*123abc`) //
	//fmt.Println(b13)
	//
	//b14 := regexp.MustCompile(`^\w\w`).MatchString(`aa+-memeda`) //\w：[a-zA-Z0-9_] ，以\w\w开头
	//fmt.Println(b14)
	//
	//b15 := regexp.MustCompile(`\d\w..`).MatchString(`123memeda`) //[0-9][a-zA-Z0-9_]任意字符任意字符
	//fmt.Println(b15)
	//
	//b16 := regexp.MustCompile(`\W`).MatchString(`b\nmemeda`) //
	//fmt.Println(b16)
	//
	//b17 := regexp.MustCompile(`\w\W`).MatchString(`a*b = 20`) //
	//fmt.Println(b17)
	//
	//b18 := regexp.MustCompile(`a[b-d]\w\d\W`).MatchString(`abc1\tdefg`) //
	//fmt.Println(b18)
	//
	//b19 := regexp.MustCompile(`a\s`).MatchString(`a bc`) //匹配空白字符。即空格
	//fmt.Println(b19)

	//数量词
	a1 := regexp.MustCompile(`\d*`).MatchString(`123`)//*是0次或多次
	fmt.Println(a1)
	a2 := regexp.MustCompile(`\d*`).MatchString(`abc123`)
	fmt.Println(a2)
	a3:= regexp.MustCompile(`\d*`).MatchString(``)
	fmt.Println(a3)
	a4 := regexp.MustCompile(`a[b-f]*\d\d[xy]*`).MatchString(`abbbb12x`)
	fmt.Println(a4)
}
