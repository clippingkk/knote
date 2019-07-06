package main

import (
	"fmt"

	"github.com/clippingkk/knote/conf"
)

func main() {
	fmt.Printf("%+v\n", conf.GetConfig())
}