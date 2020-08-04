package main

import (
	"github.com/tplish-itpm/demo/routers"
)

func main() {
	_ = routers.InitRouter().Run(":7777")
}
