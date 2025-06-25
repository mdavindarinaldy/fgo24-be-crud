package main

import (
	"backend2/router"
	"backend2/utils"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := utils.DBConnect()
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	r := gin.Default()
	router.CombineRouter(r)
	r.Run(fmt.Sprintf(":%s", os.Getenv("APP_PORT")))
}
