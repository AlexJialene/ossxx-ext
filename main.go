package main

import (
	"github.com/gogf/gf/frame/g"
	_ "ossxx-ext/boot"
	_ "ossxx-ext/route"
)

const (
	key      = ""
	secret   = ""
	bucket   = ""
	endpoint = ""
)

func main() {
	g.Server().Run()
}
