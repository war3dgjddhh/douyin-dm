package main

//	"github.com/gin-gonic/gin"
import (
	"fmt"

	"github.com/war3dgjddhh/douyin-dm/repository"
	"github.com/gin-gonic/gin"
)

func main() {
	err := repository.Init()
	if err != nil {
		fmt.Printf("%s", err)
	}

	//	go service.RunMessageServer()

	r := gin.Default()

	initRouter(r)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
