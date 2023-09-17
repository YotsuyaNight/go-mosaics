package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
	"path"
	"runtime"

	"mosaics/img"
	"mosaics/usage_map"
	"mosaics/utils"
)

type icon struct {
	Icon  *img.Img
	Color [3]uint16
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

func parseIconData(iconPaths []string) map[int]icon {
	dataMap := make(map[int]icon)
	for i, iconPath := range iconPaths {
		iconImg := utils.OpenImage(iconPath)
		dataMap[i] = icon{iconImg, iconImg.AvgAreaColor(0, 0, iconImg.Width, iconImg.Height)}
	}
	return dataMap
}

func overwriteImageRange(output *image.RGBA, input *img.Img, coords image.Rectangle) {
	x0, y0, x1, y1 := coords.Min.X, coords.Min.Y, coords.Max.X, coords.Max.Y
	for y := 0; y < y1-y0; y++ {
		for x := 0; x < x1-x0; x++ {
			r, g, b := (*input).Pixels[y][x][0], (*input).Pixels[y][x][1], (*input).Pixels[y][x][2]
			output.SetRGBA64(x0+x, y0+y, color.RGBA64{r, g, b, ^uint16(0)})
		}
	}
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

	xBlkCount, yBlkCount := inputImg.Width/blk, inputImg.Height/blk
	outputImg := image.NewRGBA(image.Rect(0, 0, xBlkCount*iconBlk, yBlkCount*iconBlk))
	usedIndices := usage_map.NewUsageMap(xBlkCount, yBlkCount)

	for y := 0; y < yBlkCount; y++ {
		for x := 0; x < xBlkCount; x++ {
			avgColor := inputImg.AvgAreaColor(x*blk, y*blk, (x+1)*blk, (y+1)*blk)
			bestMatchIndex := 0
			bestMatchDiff := ^uint64(0)

			for i, icon := range iconMap {
				sr, sg, sb := avgColor[0], avgColor[1], avgColor[2]
				ir, ig, ib := icon.Color[0], icon.Color[1], icon.Color[2]
				diff := uint64(utils.AbsDiff(sr, ir)) + uint64(utils.AbsDiff(sg, ig)) + uint64(utils.AbsDiff(sb, ib))
				if diff < bestMatchDiff && usedIndices.CanPut(x, y, 10, i) {
					bestMatchIndex = i
					bestMatchDiff = diff
					usedIndices.Map[y][x] = i
				}
			}

			outputRange := image.Rect(x*iconBlk, y*iconBlk, (x+1)*iconBlk, (y+1)*iconBlk)
			overwriteImageRange(outputImg, iconMap[bestMatchIndex].Icon, outputRange)
		}

		ms := runtime.MemStats{}
		runtime.ReadMemStats(&ms)
		log.Printf("Heap in use: %d; stack in use: %d", ms.HeapAlloc, ms.StackSys)

		// runtime.GC()
	}

	png.Encode(outputFile, outputImg)
}
