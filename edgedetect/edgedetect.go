package edgedetect

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"

	"mosaics/utils"
)

func maxVal(a, b, c uint32) uint32 {
	if a >= b && a >= c {
		return a
	}
	if b >= a && b >= c {
		return b
	}
	return c
}

func process(inputImage *image.Image) *image.RGBA {
	outputImage := image.NewRGBA((*inputImage).Bounds())

	for x := 0; x < (*inputImage).Bounds().Dx(); x++ {
		for y := 0; y < (*inputImage).Bounds().Dy()-1; y++ {
			r, g, b, _ := (*inputImage).At(x, y).RGBA()
			nr, ng, nb, _ := (*inputImage).At(x, y+1).RGBA()
			avg := uint16(maxVal(utils.AbsDiff(r, nr), utils.AbsDiff(g, ng), utils.AbsDiff(b, nb)) * 2)
			outputImage.Set(x, y, color.NRGBA64{avg, avg, avg, 1<<16 - 1})
		}
	}

	return outputImage
}

func EdgeDetect(input string, output string) {
	inputFile, err := os.Open(input)
	if err != nil {
		log.Fatal("Could not open input file: ", err)
	}

	inputImage, _, err := image.Decode(inputFile)
	if err != nil {
		log.Fatal("Could not decode input image: ", err)
	}

	outputFile, err := os.Create(output)
	if err != nil {
		log.Fatal("Could not create output file: ", err)
	}

	outputImage := process(&inputImage)

	png.Encode(outputFile, outputImage)
	log.Println("Successfully converted image!")
}
