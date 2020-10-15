package main
import (
	cache "./gccode"
	"fmt"
)
func main() {
	table := cache.Cache("lwh")
	object := cache.NewCacheObject("fuck", "you", 0)
	table.Add(object)
	get := table.Get("fuck")
	fmt.Println(get.Key())
	fmt.Println(get.Value())

}