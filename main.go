package main

import (
	"log"
	"vote/router"
)

func init() {
	log.SetFlags(log.Llongfile | log.Ldate)
}

func main() {
	g := router.Router()
	if err := g.Run(":8080"); err != nil {
		log.Fatalln(err)
	}
}
