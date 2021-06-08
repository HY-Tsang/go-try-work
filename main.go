package main

import (
	_ "try-work/boot"
	_ "try-work/router"

	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}
