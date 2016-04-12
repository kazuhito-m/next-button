package main

import (
	"fmt"
	"github.com/bitly/go-simplejson"
)

func main() {
	fmt.Println("Hello World!")
	// 文字列json化の例
	json := simplejson.New()
	json.Set("message", "Hello, World!")
	b, _ := json.EncodePretty()
	fmt.Printf("%s\n", b)
}
