package splitter

import (
	"fmt"
	"image"
	"image/png"
	"log"
	"os"
)

func process(grid *image.Image) {
	blk := 64
	index := 1

	for x := 0; x < (*grid).Bounds().Dx()/blk; x++ {
		for y := 0; y < (*grid).Bounds().Dy()/blk; y++ {
			outputFile, err := os.Create(fmt.Sprintf("icons/icon-%d.png", index))
			if err != nil {
				log.Fatal("Could not create output file: ", err)
			}
			outputImage := image.NewRGBA(image.Rect(0, 0, blk-1, blk-1))
			for i := 0; i < blk; i++ {
				for j := 0; j < blk; j++ {
					outputImage.Set(i, j, (*grid).At(x*blk+i, y*blk+j))
				}
			}
			png.Encode(outputFile, outputImage)
			index++
		}
	}
}

func Splitter(input string) {
	inputFile, err := os.Open(input)
	if err != nil {
		log.Fatal("Could not open input file: ", err)
	}

	inputImage, _, err := image.Decode(inputFile)
	if err != nil {
		log.Fatal("Could not decode input image: ", err)
	}

	process(&inputImage)

	log.Println("Successfully converted image!")
}
