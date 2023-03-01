package main

import (
	"gindiary/model"
	"gindiary/routes"
)

func main() {
	model.InitDb()
	routes.InitRouter()
}
