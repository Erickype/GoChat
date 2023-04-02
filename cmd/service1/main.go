package main

import (
	"log"
	"time"
)

func main() {
	for {
		time.Sleep(time.Second)
		log.Println("HElLO WORLD from service 1!")
	}
}
