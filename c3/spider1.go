package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

func main()  {
	url := "http://www.zhenai.com/zhenghun"
	resp, err := http.Get(url)
	if err != nil{
		log.Println(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK{
		fmt.Println("Error: status code:", resp.StatusCode)
		return
	}
	all, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		fmt.Println(err)
	}
	//fmt.Printf("%s\n", all)

	//re := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)" [^>]*>([^<]+)</a>`)
	//content := re.FindAllSubmatch(all, -1)
	//fmt.Println(len(content))
	//for _,c := range content{
	//	fmt.Printf("City \t%s\t,URL\t%s", c[2],c[1])
	//	fmt.Println()
	//}

	re:= regexp.MustCompile(`<script>window.__INITIAL_STATE__=(.+);\(function`)
	content := re.FindSubmatch(all)
	fmt.Printf("%s",content[1])
}
