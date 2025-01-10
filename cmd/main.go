package main

import (
	"flag"
	"fmt"
	"navpage/routes"
)

var (
	Port = flag.Int("port", 8080, "服务端口")
)

func main() {
	flag.Parse()

	app := routes.Routes()
	app.Listen(fmt.Sprintf(":%d", *Port))
}
