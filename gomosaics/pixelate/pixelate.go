package pixelate

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

func process(inputImage *image.Image, blk int) *image.RGBA {
	outputImage := image.NewRGBA((*inputImage).Bounds())

	for x := 0; x < (*inputImage).Bounds().Dx()/blk; x++ {
		for y := 0; y < (*inputImage).Bounds().Dy()/blk; y++ {
			var tr, tg, tb uint64
			for i := 0; i < blk; i++ {
				for j := 0; j < blk; j++ {
					pixel := (*inputImage).At(x*blk+i, y*blk+j)
					r, g, b, _ := pixel.RGBA()
					// log.Printf("Decoded a pixel: %b / %b / %b", r, g, b)
					tr += uint64(r)
					tg += uint64(g)
					tb += uint64(b)
				}
			}
			tr = tr / uint64(blk*blk)
			tg = tg / uint64(blk*blk)
			tb = tb / uint64(blk*blk)
			for i := 0; i < blk; i++ {
				for j := 0; j < blk; j++ {
					newPixel := color.NRGBA64{uint16(tr), uint16(tg), uint16(tb), 1<<16 - 1}
					outputImage.Set(x*blk+i, y*blk+j, newPixel)
				}
			}
		}
	}

	return outputImage
}

func Pixelate(input string, output string, blk int) {
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

	outputImage := process(&inputImage, blk)

	png.Encode(outputFile, outputImage)
	log.Println("Successfully converted image!")
}
