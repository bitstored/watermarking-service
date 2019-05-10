package steganography

import (
	"image"
	"image/draw"
)

func imageToRGBA(src image.Image) *image.RGBA {
	bounds := src.Bounds()

	var m *image.RGBA
	var width, height int

	width = bounds.Dx()
	height = bounds.Dy()

	m = image.NewRGBA(image.Rect(0, 0, width, height))

	draw.Draw(m, m.Bounds(), src, bounds.Min, draw.Src)
	return m
}

func bitCount(x, y int) int {
	return (x-1)*2 + (y-1)*2
}

func getCoordinatesFromIndex(dx int, dy int, indx int) (int, int) {
	if indx > bitCount(dx, dy) {
		return -1, -1
	}
	rounds := indx / 4
	var x int
	var y int
	if rounds >= dx-2 {
		rounds = (indx - 2*(dx-2)) / 2
		side := indx % 2
		x = (dx - 1) - (dx-1)*side
		y = rounds + 1
	} else if rounds >= dy-2 {
		rounds = (indx - 2*(dy-2)) / 2
		side := indx % 2
		y = (dy - 1) - (dy-1)*side
		x = rounds + 1
	} else {
		side := indx % 4
		switch side {
		case 0:
			x = 0
			y = indx/4 + 1
		case 1:
			x = dx - 1
			y = indx/4 + 1
		case 2:
			y = 0
			x = indx/4 + 1
		case 3:
			y = dy - 1
			x = indx/4 + 1
		}
	}
	return x, y
}

func getIndexFromCoordinates(dx int, dy int, x int, y int) (indx int, offset int) {
	if dx < x || dy < y {
		return -1, -1
	}

	if x == 0 {
		indx += 2 * (y - 1)
		indx += 2 * min(dx-2, y-1)
	} else if x == dx-1 {
		indx += 2 * (y - 1)
		indx += 2 * min(dx-2, y-1)
		indx += 1
	} else if y == 0 {
		indx += 2 * (y - 1)
		indx += 2 * min(dy-2, x-1)
	} else if y == dy-1 {
		indx += 2 * (y - 1)
		indx += 2 * min(dy-2, x-1)
		indx += 1
	}

	offset = indx % 8
	indx /= 8
	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func setBit(a *uint32, b uint32, cnt uint) {
	mask := b & (1 << cnt)
	if b&mask == 0 {
		*a = (*a) & (255 - 1)
	} else {
		*a = (*a) | 1
	}
}
