package main

import (
	"github.com/kashimashino/hackweek/model"
	"github.com/kashimashino/hackweek/routes"
)

func main() {
	model.InitDb()

	routes.InitRouter()
}


