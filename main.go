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
	// image := utils.OpenImage("bg3.png")
	// for i := 0; i < 20000000000; i++ {
	// 	pixel := image.At(512, 512)
	// 	pixel = pixel
	// }

	log.Println("Finished! Press <Enter> to finish...")
	fmt.Scanln()
}
