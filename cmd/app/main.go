package main

import (
	"log"
	"wiltonassessment/internal/app"
)

const cfgFile = "configs/srvconfig.yml"

func main() {
	if err := app.Run(cfgFile); err != nil {
		log.Fatal(err)
	}
}
