package cave

import (
	"image"
	"image/draw"
	"image/png"
	"log"
	"math"
	"math/rand"
	"newton-pcg/core"
	"newton-pcg/perlin"
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
	grassPath   = "assets/Grass.png"

	treePath  = "assets/Tree_SpriteSheet_Outlined.png"
	imageSize = 8
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
var tree = getImageFromFilePath(treePath)
var grass = getImageFromFilePath(grassPath)

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

func DrawWorld(m core.Field) {
	//m = smooth(m)

	canvas := image.NewRGBA(image.Rect(0, 0, m.W*imageSize, m.H*imageSize))
	r := image.Rect(0, 0, 16, 16)

	core.MeshCB4G(m.W, m.H, func(p core.P) {
		offset := image.Pt(p.X*imageSize, p.Y*imageSize)
		pos := r.Add(offset)
		v := m.At(p)
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
		case 10:
			brush = grass
		default:
			brush = white
		}
		draw.Draw(canvas, pos, brush, image.Point{2, 2}, draw.Src)
	})

	core.SaveImage(canvas)
}

func SurfaceMask(w, h, level int) core.Field {
	img := make([][]int, h)
	for i := 0; i < h; i++ {
		img[i] = make([]int, w)
	}
	sinsum := randomSin(w, level)
	for x := 0; x < w; x++ {
		lh := sinsum(x)
		for i := level + lh; i < h; i++ {
			img[i][x] = 1
		}
	}
	return core.Field{
		W: w,
		H: h,
		F: img,
	}
}

func Mask(W, H int, a, b float64, seed int64) core.Field {
	perl := perlin.NewPerlin(a, b, 3, seed)

	f := core.NewField(W, H)

	for p := range core.Mesh(W, H) {
		pv := int(perl.Noise2D(float64(p.X)/float64(W), float64(p.Y)/float64(H)) * 255)
		var v int
		if pv > 100 || pv < 60 {
			v = 1
		}
		f.Set(p, v)
	}
	return f
}

func randomSin(level, w int) func(int) int {
	var sins []func(x float64) float64
	for i := 0; i < 4; i++ {
		w := float64(1+rand.Intn(8)) / 10
		A := 5. * float64(1+rand.Intn(5))
		sins = append(sins, func(x float64) float64 {
			return A * math.Sin(x/w)
		})
	}
	return func(x int) int {
		r := 0
		for _, f := range sins {
			r += int(f(float64(x) / float64(w)))
		}
		if level+r < 0 {
			return 0
		}
		return r
	}
}

func randomTree() (offset, size core.P) {
	x := rand.Intn(10)
	y := x % 2
	treeW := tree.Bounds().Max.X / 5
	treeH := tree.Bounds().Max.Y / 2
	return core.P{
			X: x * treeW,
			Y: y * treeH,
		}, core.P{
			X: treeW,
			Y: treeH,
		}
}

func GroundLayer(w, h, level int) core.Field {
	img := make([][]int, h)
	for i := 0; i < h; i++ {
		img[i] = make([]int, w)
	}
	sinsum := randomSin(w, level)
	for x := 0; x < w; x++ {
		lh := sinsum(x)
		for i := level + lh; i > 0; i-- {
			img[i][x] = 1
		}
	}
	return core.Field{
		W: w,
		H: h,
		F: img,
	}
}

func GrassLayer(img core.Field, from, to int) core.Field {
	w := img.W
	h := img.H
	for p := range core.Mesh(w, h) {
		if p.Y == 0 {
			continue
		}
		if img.At(p) != from {
			continue
		}
		if img.At(core.P{X: p.X, Y: p.Y - 1}) == 0 {
			img.Set(p, to)
		}
	}
	return img
}
