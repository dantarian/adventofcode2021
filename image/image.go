package image

type point struct {
	x, y int
}

type Image struct {
	pixels               map[point]bool
	topLeft, bottomRight point
	externalValue        bool
}

func NewImage(data []string) *Image {
	image := Image{pixels: make(map[point]bool), topLeft: point{0, 0}}
	maxX, maxY := 0, 0
	for y, line := range data {
		for x, char := range line {
			image.pixels[point{x, y}] = char == '#'
			maxX = x
		}
		maxY = y
	}
	image.bottomRight = point{maxX, maxY}

	return &image
}

func (i *Image) Enhance(alg Algorithm) *Image {
	flipExternal := false
	if alg[0] {
		flipExternal = true
	}
	newImage := Image{
		pixels:        make(map[point]bool),
		topLeft:       point{i.topLeft.x - 1, i.topLeft.y - 1},
		bottomRight:   point{i.bottomRight.x + 1, i.bottomRight.y + 1},
		externalValue: (i.externalValue && !flipExternal) || (flipExternal && !i.externalValue),
	}

	for x := newImage.topLeft.x; x <= newImage.bottomRight.x; x++ {
		for y := newImage.topLeft.y; y <= newImage.bottomRight.y; y++ {
			p := point{x, y}
			neighbours := p.Neighbours()
			value := 0
			for _, neighbour := range neighbours {
				value *= 2
				if i.Value(neighbour) {
					value += 1
				}
			}
			newImage.pixels[p] = alg[value]
		}
	}

	return &newImage
}

func (p point) Neighbours() []point {
	return []point{
		{p.x - 1, p.y - 1},
		{p.x, p.y - 1},
		{p.x + 1, p.y - 1},
		{p.x - 1, p.y},
		{p.x, p.y},
		{p.x + 1, p.y},
		{p.x - 1, p.y + 1},
		{p.x, p.y + 1},
		{p.x + 1, p.y + 1},
	}
}

func (i *Image) Value(p point) bool {
	if p.x < i.topLeft.x || p.x > i.bottomRight.x || p.y < i.topLeft.y || p.y > i.bottomRight.y {
		return i.externalValue
	}

	return i.pixels[p]
}

func (i *Image) LitPixelCount() int {
	count := 0
	for _, pixel := range i.pixels {
		if pixel {
			count++
		}
	}
	return count
}
