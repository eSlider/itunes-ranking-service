package main

import "github.com/eSlider/itunes-ranking-service/api"

func main() {
	api.NewService("swagger.yml").Start()
}
