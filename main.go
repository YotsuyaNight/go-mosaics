package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	Mosaicate("bg3.png", "icons_small")

	log.Println("Finished! Press <Enter> to finish...")
	fmt.Scanln()
}
