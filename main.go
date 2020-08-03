package main

import (
	"fmt"
	"io/ioutil"
	"time"

	"github.com/by12380/Autocomplete/routers"
	"github.com/by12380/Autocomplete/services/trie"
	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

func main() {
	start := time.Now()

	r := gin.Default()

	routers.InitAutocomplete(r.Group("/autocomplete"))

	bytes, err := ioutil.ReadFile("./assets/input.json")

	if err != nil {
		println(err)
	}

	result := gjson.GetBytes(bytes, "items")
	// println(result.Raw)

	t := trie.GetInstance()
	for _, item := range result.Array() {
		textItem := trie.TextItem{
			Value:  item.Get("value").String(),
			Weight: int(item.Get("weight").Int()),
		}
		t.Add(&textItem)
	}

	r.Run()

	duration := time.Since(start)
	fmt.Println(duration)
}
