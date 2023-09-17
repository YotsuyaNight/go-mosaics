package utils

import (
	"image"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
)

func OpenImage(input string) image.Image {
	inputFile, err := os.Open(input)
	if err != nil {
		log.Fatal("Cannot open input file: ", err)
	}
	defer inputFile.Close()
	inputImg, _, err := image.Decode(inputFile)
	if err != nil {
		log.Fatal("Cannot decode input image: ", err)
	}
	return inputImg
}

func AbsDiff(x, y uint32) uint32 {
	if x < y {
		return y - x
	}
	return x - y
}
