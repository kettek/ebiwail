package main

import (
	"fmt"
	"image/color"
	"math/rand"
	"syscall/js"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Rect struct {
	x, y, width, height int
	color               color.NRGBA
}

func MakeRandomRect(maxW, maxH int) Rect {
	r := uint8(rand.Intn(255))
	g := uint8(rand.Intn(255))
	b := uint8(rand.Intn(255))

	if maxW == 0 {
		maxW = 320
	}
	if maxH == 0 {
		maxH = 240
	}

	return Rect{
		x:      rand.Intn(maxW),
		y:      rand.Intn(maxH),
		width:  rand.Intn(10) + 10,
		height: rand.Intn(10) + 10,
		color:  color.NRGBA{r, g, b, 255},
	}
}

type Game struct {
	w, h  int
	rects []Rect
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, r := range g.rects {
		vector.DrawFilledRect(screen, float32(r.x), float32(r.y), float32(r.width), float32(r.height), r.color, true)
	}
	ebitenutil.DebugPrint(screen, fmt.Sprintf("%dx%d", g.w, g.h))
}

func (g *Game) Layout(ow, oh int) (int, int) {
	g.w = ow
	g.h = oh
	return ow, oh
}

func main() {
	var game Game

	js.Global().Get("window").Call("addEventListener", "message", js.FuncOf(func(this js.Value, p []js.Value) any {
		event := p[0].Get("data").String()
		switch event {
		case "button:random":
			game.rects = append(game.rects, MakeRandomRect(game.w, game.h))
		case "button:clear":
			game.rects = []Rect{}
		default:
			fmt.Println("Unknown event:", event)
		}
		return nil
	}))

	ebiten.SetWindowSize(320, 240)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&game); err != nil {
		panic(err)
	}
}
