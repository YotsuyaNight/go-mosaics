package img

import (
	"image"
)

type Img struct {
	Pixels        [][][3]uint16
	Width, Height int
}

func New(w, h int) *Img {
	pixels := make([][][3]uint16, h)

	for y := 0; y < h; y++ {
		pixels[y] = make([][3]uint16, w)
	}

	newImg := Img{
		Pixels: pixels,
		Width:  w,
		Height: h,
	}
	return &newImg
}

func FromBuiltin(builtinImg image.RGBA64Image) *Img {
	w, h := builtinImg.Bounds().Dx(), builtinImg.Bounds().Dy()
	pixels := make([][][3]uint16, h)

	for y := 0; y < h; y++ {
		pixels[y] = make([][3]uint16, w)
		for x := 0; x < w; x++ {
			pixel := builtinImg.RGBA64At(x, y)
			pixels[y][x][0] = pixel.R
			pixels[y][x][1] = pixel.G
			pixels[y][x][2] = pixel.B
		}
	}

	newImg := Img{
		Pixels: pixels,
		Width:  w,
		Height: h,
	}
	return &newImg
}

func (img *Img) AvgAreaColor(x0, y0, x1, y1 int) [3]uint16 {
	total := uint64((x1 - x0) * (y1 - y0))
	var r, g, b uint64
	for y := y0; y < y1; y++ {
		for x := x0; x < x1; x++ {
			r += uint64(img.Pixels[y][x][0])
			g += uint64(img.Pixels[y][x][1])
			b += uint64(img.Pixels[y][x][2])
		}
	}
	return [3]uint16{uint16(r / total), uint16(g / total), uint16(b / total)}
}
