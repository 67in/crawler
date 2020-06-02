package main

import (
	"regexp"
	"fmt"
)

//func main()  {
//	const text = "My email is liuqin@163.com@qq.com"
//	//re := regexp.MustCompile(`[a-zA-Z0-9]+@[a-zA-Z0-9.]+\.[a-zA-Z0-9]+`)
//	//match := re.FindString(text)
//
//	re, err := regexp.Compile("liuqin@163.com")
//	if err != nil{
//		println(err)
//	}
//	match := re.FindString(text)
//	fmt.Println(match)
//}

func main() {
	const text = `My email is liuqin@163.com
                  email is liuzhenkun@qq.com
                  email is liushukai@qq.com`
	re ,err := regexp.Compile(`[a-zA-Z0-9]+@[a-zA-Z0-9.]+\.[a-zA-Z0-9]+`)
	if err != nil{
		println(err)
	}
	match := re.FindAllString(text, -1)
	fmt.Println(match)
}
