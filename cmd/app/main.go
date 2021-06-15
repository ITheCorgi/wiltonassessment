package main

import (
	"log"
	"wiltonassessment/internal/app"
)

const cfgFile = "configs/serverconfig.yml"

func main() {
	if err := app.Run(cfgFile); err != nil {
		log.Println(err.Error())
	}
}
