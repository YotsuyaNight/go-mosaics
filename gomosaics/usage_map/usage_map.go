package usage_map

type UsageMap struct {
	Width, Height int
	Map           [][]int
}

func NewUsageMap(width, height int) *UsageMap {
	mapValue := make([][]int, height)
	for i := 0; i < height; i++ {
		mapValue[i] = make([]int, width)
	}

	return &UsageMap{
		Width:  width,
		Height: height,
		Map:    mapValue,
	}
}

func (um *UsageMap) CanPut(x, y, radius, value int) bool {
	for i := max(0, x-radius); i < min(x+radius, um.Width); i++ {
		for j := max(0, y-radius); j < min(y+radius, um.Height); j++ {
			if um.Map[j][i] == value {
				return false
			}
		}
	}
	return true
}
