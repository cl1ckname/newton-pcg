package cave

import (
	"image"
	"image/draw"
	"image/png"
	"log"
	"newton-pcg/core"
	"os"
)

const (
	dirtPath    = "assets/Dirt_Block.png"
	stonePath   = "assets/Stone_Block.png"
	ironOrePath = "assets/Iron_Ore_(placed).png"
	goldOrePath = "assets/Gold_Ore_(placed).png"
	blackPath   = "assets/black.png"
	imageSize   = 16
)

func getImageFromFilePath(filePath string) image.Image {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	img, err := png.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	return img
}

func GenerateCave(m [][]int) {
	m = smooth(m)
	h := len(m)
	w := len(m[0])
	stone := getImageFromFilePath(stonePath)
	dirt := getImageFromFilePath(dirtPath)
	iron := getImageFromFilePath(ironOrePath)
	gold := getImageFromFilePath(goldOrePath)
	black := getImageFromFilePath(blackPath)

	canvas := image.NewRGBA(image.Rect(0, 0, w*imageSize, h*imageSize))
	r := image.Rect(0, 0, 16, 16)
	for y, row := range m {
		for x, v := range row {
			offset := image.Pt(x*imageSize, y*imageSize)
			pos := r.Add(offset)
			var brush image.Image
			switch v {
			case 0:
				brush = stone
			case 1:
				brush = dirt
			case 2:
				brush = iron
			case 3:
				brush = gold
			default:
				brush = black
			}
			draw.Draw(canvas, pos, brush, image.Point{}, draw.Src)
		}
	}
	core.SaveImage(canvas)
}

func smooth(m [][]int) [][]int {
	h := len(m)
	w := len(m[0])
	for p := range core.Mesh(w-2, h-2) {
		x := p.X + 1
		y := p.Y + 1
		s := m[y][x+1] + m[y][x-1] + m[y+1][x] + m[y-1][x]
		m[y][x] = s / 4
	}
	return m
}
