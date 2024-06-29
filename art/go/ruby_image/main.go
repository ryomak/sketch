package main

import (
	"image/color"
	"syscall/js"
	"math"
	"github.com/ryomak/sketch/art"
	"github.com/go-p5/p5"
)

func main() {
	art.Do("go_ruby_image",Art{})
}

var _ art.Interface = Art{}
type Art struct{}

func (Art) Title() any {
	return "幾何学"
}

func (Art) Description() any {
	return "幾何学aaaa"
}

func (Art) Generate(this js.Value, args []js.Value) any {
	p5.Run(setup, draw)
	return nil
}

var (
	angle float64
)

func setup() {
	p5.Canvas(400, 400)
}


func draw() {
	// Ruby red background
	p5.Background(color.RGBA{R: 220, G: 20, B: 60, A: 255})

	// Draw a simplified Ruby "gem"
	p5.Push()
	p5.Translate(200, 200)
	p5.Rotate(angle)
	p5.Fill(color.RGBA{R: 255, G: 0, B: 0, A: 255})
	p5.Rect(-75, -75, 150, 150)
	p5.Pop()

	// Draw floating code snippets
	p5.Fill(color.White)
	p5.TextSize(16)
	p5.Text("def ruby_method", 50, 50)
	p5.Text("puts 'Hello, Ruby!'", 250, 200)
	p5.Text("end", 100, 350)

	// Draw rotating dots
	for i := 0; i < 8; i++ {
		x := 200 + math.Cos(angle+float64(i)*math.Pi/4)*150
		y := 200 + math.Sin(angle+float64(i)*math.Pi/4)*150
		p5.Fill(color.RGBA{R: 255, G: 255, B: 255, A: 200})
		p5.Ellipse(x, y, 20, 20)
	}

	angle += 0.02
}


