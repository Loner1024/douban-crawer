package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

const url string = "https://book.douban.com"

func main()  {
	client:=&http.Client{
	
	}
	req,err:=http.NewRequest("GET",url+"/tag/程序",nil)
	req.Header = map[string][]string{
		"User-Agent": []string{"Mozilla/5.0 (Macintosh; Intel Mac OS X 11_2_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.146 Safari/537.36"},
		}
	if err!=nil{
		panic(err)
	}
	resp,err:=client.Do(req)
	if err!=nil{
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode !=http.StatusOK{
		log.Printf("Error: status code %d\n",resp.StatusCode)
		return
	}
	all,err:=ioutil.ReadAll(resp.Body)
	// fmt.Printf("%s\n",all)
	printTagList(all)
}


func printTagList(contents []byte){
	re:=regexp.MustCompile(`<a href="(/tag/[^>]*)">([^<]*)</a>`)
	match:=re.FindAllStringSubmatch(string(contents),-1)
	for _,v :=range match{
		fmt.Printf("url: %s%s,tag:%s\n",url,v[1],v[2])
	}}