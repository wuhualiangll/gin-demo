package main

import "gin-demo/router"

func main() {
	r := router.SetupRouter()
	r.Run(":8080")
}
