package main

import (
	_ "fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	
	router.GET("/", func(c* gin.Context) {
		c.String(200, "Wsg, boy!")
    	c.String(200, "My name is Monteiro and im just a student.")
	})

	router.GET("/interests", func(c* gin.Context) {
		c.String(200, "I love programming. I code since 2023 when i was just 14 y/o.")
	})

	router.GET("/friends", func(c* gin.Context) {
		c.String(200, "Theres some friends, people i really enjoy talking.")
		c.String(200, "Wender, m2hcz, X0zy, kiino")
	})

	router.Run()
}
