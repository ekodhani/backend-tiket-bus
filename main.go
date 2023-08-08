package main

import (
	"backend-pertama/config"
)

func main() {
	router := config.SetupRouter()
	port := config.Port()

	router.Run(port)
}
