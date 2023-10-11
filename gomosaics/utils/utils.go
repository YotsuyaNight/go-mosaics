package utils

import (
	"image"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"

	"gomosaics/img"
)

func OpenImage(input string) (*img.Img, error) {
	inputFile, err := os.Open(input)
	if err != nil {
		log.Println("Cannot open input file: ", err)
		return nil, err
	}
	defer inputFile.Close()
	inputImg, _, err := image.Decode(inputFile)
	if err != nil {
		log.Println("Cannot decode input image: ", err)
		return nil, err
	}
	return img.FromBuiltin(inputImg.(image.RGBA64Image)), nil
}

func AbsDiff[T uint8 | uint16](x, y T) T {
	if x < y {
		return y - x
	}
	return x - y
}
