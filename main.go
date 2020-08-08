package main

import (
	"io/ioutil"

	"github.com/by12380/Autocomplete/routers"
	"github.com/by12380/Autocomplete/services/trie"
	"github.com/caarlos0/env/v6"
	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

type config struct {
	AssetsPath string `env:"ASSETS_PATH"`
}

func main() {
	cfg := config{}
	env.Parse(&cfg)

	r := gin.Default()

	routers.InitAutocomplete(r.Group("/autocomplete"))

	bytes, err := ioutil.ReadFile(cfg.AssetsPath + "/input.json")

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
}
