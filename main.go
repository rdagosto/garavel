package main

import (
	"fmt"
	"garavel/internal/configs"
	"garavel/internal/routes"
)

func main() {
	fmt.Println("#### Garavel ####")

	routes.HandleRequest()
	routes.R.Run(":" + configs.Env("APP_PORT", "8080"))
}
