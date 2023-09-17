package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
	"path"
	"runtime"

	"mosaics/usage_map"
	"mosaics/utils"
)

type icon struct {
	Icon image.Image
	color.Color
}

func scanIconsDir(dir string) []string {
	entries, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal("Could not open icon dir: ", err)
	}

	var files []string
	for _, entry := range entries {
		if !entry.IsDir() {
			files = append(files, path.Join(dir, entry.Name()))
		}
	}

	return files
}

func avgAreaColor(img image.Image, area image.Rectangle) color.Color {
	x0, y0, x1, y1 := area.Min.X, area.Min.Y, area.Max.X, area.Max.Y
	r, g, b, _ := img.At(x0, y0).RGBA()
	ar, ag, ab := float32(r), float32(g), float32(b)
	for y := y0; y <= y1; y++ {
		for x := x0 + 1; x <= x1; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			t := float32((y-y0)*(x-x0) + (x - x0) + 1)
			ar = ar + float32(1.0/t)*(float32(r)-ar)
			ag = ag + float32(1.0/t)*(float32(g)-ag)
			ab = ab + float32(1.0/t)*(float32(b)-ab)
		}
	}
	return color.NRGBA64{uint16(ar), uint16(ag), uint16(ab), 1<<16 - 1}
}

func parseIconData(iconPaths []string) map[int]icon {
	dataMap := make(map[int]icon)
	for i, iconPath := range iconPaths {
		iconImg := utils.OpenImage(iconPath)
		dataMap[i] = icon{iconImg, avgAreaColor(iconImg, iconImg.Bounds())}
	}
	return dataMap
}

func overwriteImageRange(output *image.RGBA, input image.Image, coords image.Rectangle) {
	x0, y0, x1, y1 := coords.Min.X, coords.Min.Y, coords.Max.X, coords.Max.Y
	for y := 0; y < y1-y0; y++ {
		for x := 0; x < x1-x0; x++ {
			output.Set(x0+x, y0+y, input.At(x, y))
		}
	}
}

func blkRect(x, y, blk int) image.Rectangle {
	return image.Rect(x*blk, y*blk, (x+1)*blk, (y+1)*blk)
}

func Mosaicate(input string, iconDir string) {
	log.Println("Running mosaicate with", input, iconDir)
	blk := 2
	iconBlk := 16
	iconMap := parseIconData(scanIconsDir(iconDir))
	inputImg := utils.OpenImage(input)

	outputFile, err := os.Create("output.png")
	if err != nil {
		log.Fatal("Cannot create output file: ", err)
	}

	xBlkCount, yBlkCount := inputImg.Bounds().Dx()/blk, inputImg.Bounds().Dy()/blk
	outputImg := image.NewRGBA(image.Rect(0, 0, xBlkCount*iconBlk, yBlkCount*iconBlk))
	usedIndices := usage_map.NewUsageMap(xBlkCount, yBlkCount)

	for y := 0; y < yBlkCount; y++ {
		for x := 0; x < xBlkCount; x++ {
			avgColor := avgAreaColor(inputImg, blkRect(x, y, blk))
			bestMatchIndex := 0
			bestMatchDiff := ^uint32(0)

			for i, icon := range iconMap {
				sr, sg, sb, _ := avgColor.RGBA()
				ir, ig, ib, _ := icon.RGBA()
				diff := utils.AbsDiff(sr, ir) + utils.AbsDiff(sg, ig) + utils.AbsDiff(sb, ib)
				if diff < bestMatchDiff && usedIndices.CanPut(x, y, 10, i) {
					bestMatchIndex = i
					bestMatchDiff = diff
					usedIndices.Map[y][x] = i
				}
			}

			overwriteImageRange(outputImg, iconMap[bestMatchIndex].Icon, blkRect(x, y, iconBlk))
		}

		ms := runtime.MemStats{}
		runtime.ReadMemStats(&ms)
		log.Printf("Heap in use: %d; mallocs: %d", ms.HeapAlloc, ms.Mallocs)

		// runtime.GC()
	}

	png.Encode(outputFile, outputImg)
}
