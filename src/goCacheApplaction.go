package main

import (
	cache "./gccode"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)
/*
set table key value

 */
func main() {
	var currentTable *cache.CacheTable
	currentTable = cache.Cache("defatul")

	fmt.Println("Welcome to use letitgo-cache")
	fmt.Println("The current table is",currentTable.Name())
	fmt.Println()
	scanner :=bufio.NewScanner(os.Stdin)

	for scanner.Scan(){
		text := scanner.Text()
		if len(strings.TrimSpace(text))==0{
			continue
		}
		coms := strings.Split(text, " ")
		switch coms[0] {
		case "tables":
			tbns:=cache.Tables()
			for i,tn:=range tbns{
				fmt.Printf("|%3d|%12s|\n",i,tn)
			}
		case "select":
			if len(coms)<2{
				fmt.Println("One of the tables must be selected")
				break
			}
			tbn := coms[1]
			if len(strings.TrimSpace(tbn))==0{
				fmt.Println("One of the tables must be selected")
			}else{
				currentTable = cache.Cache(tbn)
				fmt.Println("The current table is",currentTable.Name())
			}
		case "set":{
			if len(coms)<3{
				fmt.Println("tip: set key var [exp /s]")
				break
			}
			if len(coms)>3{
				var s string= coms[3]
				parseInt, err := strconv.ParseInt(s, 10, 32)
				duration, err := time.ParseDuration(strconv.FormatInt(parseInt, 10) + "s")

				if err !=nil{
					fmt.Println(err)
				}
				currentTable.Add(cache.NewCacheObject(coms[1],coms[2],duration))
			}else{
				currentTable.Add(cache.NewCacheObject(coms[1],coms[2],0))
			}


		}
		case "get":
			if len(coms)<2{
				fmt.Println("tip: get key")
				break
			}
			get := currentTable.Get(coms[1])
			if get ==nil{
				fmt.Println(nil)
			}else {
				fmt.Println(get.Value())
			}
		case "help":
				fmt.Println("  ")
				fmt.Println("tips:")
				fmt.Println("opt:")
				fmt.Println("    - select tableName : 选择/创建库 select defatul 或  select newDb ")
				fmt.Println("    - tables  : 展示所有库 ")
				fmt.Println("    - set  : 设置当前库的缓存数据 set key value [exp],exp为存活时间(s) ")
				fmt.Println("    - get  : 获取当前库的缓存数据 get key ")
				fmt.Println("  ")

		}

	}

}
