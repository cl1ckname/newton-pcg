package cave

import (
	"image"
	"image/draw"
	"image/png"
	"log"
	"math"
	"math/rand"
	"newton-pcg/core"
	"os"
)

const (
	dirtPath    = "assets/Dirt_Block.png"
	stonePath   = "assets/Stone_Block.png"
	ironOrePath = "assets/Iron_Ore_(placed).png"
	goldOrePath = "assets/Gold_Ore_(placed).png"
	blackPath   = "assets/black.png"
	siltPath    = "assets/Silt_Block.png"
	whitePath   = "assets/White_background.png"
	marblePath  = "assets/Marble_Block.png"
	fleshPath   = "assets/Flesh_Block.png"
	mythrilPath = "assets/Mythril_Ore.png"
	imageSize   = 8
)

var stone = getImageFromFilePath(stonePath)
var dirt = getImageFromFilePath(dirtPath)
var iron = getImageFromFilePath(ironOrePath)
var gold = getImageFromFilePath(goldOrePath)
var black = getImageFromFilePath(blackPath)
var silt = getImageFromFilePath(siltPath)
var white = getImageFromFilePath(whitePath)
var marble = getImageFromFilePath(marblePath)
var flesh = getImageFromFilePath(fleshPath)
var mythril = getImageFromFilePath(mythrilPath)

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

func DrawWorld(m [][]int) {
	//m = smooth(m)
	h := len(m)
	w := len(m[0])

	canvas := image.NewRGBA(image.Rect(0, 0, w*imageSize, h*imageSize))
	r := image.Rect(0, 0, 16, 16)
	for y, row := range m {
		for x, v := range row {
			offset := image.Pt(x*imageSize, y*imageSize)
			pos := r.Add(offset)
			var brush image.Image
			switch v {
			case 0:
				brush = black
			case 1:
				brush = stone
			case 2:
				brush = dirt
			case 3:
				brush = silt
			case 4:
				brush = marble
			case 5:
				brush = iron
			case 6:
				brush = gold
			case 7:
				brush = flesh
			case 8:
				brush = mythril
			default:
				brush = white
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

func SurfaceMask(w, h, level int) [][]int {
	img := make([][]int, h)
	for i := 0; i < h; i++ {
		img[i] = make([]int, w)
	}
	var sins []func(x float64) float64
	for i := 0; i < 6; i++ {
		w := float64(1+rand.Intn(8)) / 100
		A := 5. * float64(1+rand.Intn(5))
		sins = append(sins, func(x float64) float64 {
			return A * math.Sin(x/w)
		})
	}
	sinsum := func(x int) int {
		r := 0
		for _, f := range sins {
			r += int(f(float64(x) / float64(w)))
		}
		if level+r < 0 {
			return 0
		}
		return r
	}
	for x := 0; x < w; x++ {
		lh := sinsum(x)
		for i := level + lh; i < h; i++ {
			img[i][x] = 1
		}
	}
	return img
}

func CavesMask(proto [][]int, thresh int) [][]int {
	h := len(proto)
	w := len(proto[0])
	img := make([][]int, h)
	for i := 0; i < h; i++ {
		img[i] = make([]int, w)
	}
	for p := range core.Mesh(w, h) {
		if proto[p.Y][p.X] < thresh {
			img[p.Y][p.X] = 1
		}
	}
	return img
}
