package main

import (
	router "ginchat/router"
	"ginchat/utils"
)

func main() {
	utils.InitConfig()
	utils.InitMySQL()
	r := router.Router()
	r.Run("localhost:8081") // listen and serve on
}
