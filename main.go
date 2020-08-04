package main

import (
	"github.com/tplish-itpm/demo/models"
	"github.com/tplish-itpm/demo/routers"
)

func main() {
	models.InitDB()
	routers.InitRouter()
}
