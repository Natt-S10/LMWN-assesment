package main

import (
	"fmt"
	_ "fmt"

	"github.com/Natt-S10/LMWN-assesment/config"
	"github.com/Natt-S10/LMWN-assesment/internal/routes"
)

func main() {

	router := routes.SetupRouter()
	fmt.Println("port: ", config.SERVER_PORT)
	router.Run(":" + config.SERVER_PORT)
}
