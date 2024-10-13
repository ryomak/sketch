package main

import (
	"syscall/js"
	"math"

	"github.com/ryomak/sketch/art"
	"github.com/ryomak/p5go"
)

func main() {
	art.Do("stylish_ruby", Art{})
}

var _ art.Interface = Art{}
type Art struct{}

func (Art) Title() any {
	return "Stylish Ruby"
}

func (Art) Description() any {
	return "An elegant representation of Ruby using geometric animations"
}

func (Art) Generate(this js.Value, args []js.Value) any {
	p5go.Execute(args[0].String(),
		p5go.WithSetup(setup),
		p5go.WithDraw(draw),
	)
	return nil
}

var angle float64

func setup(p *p5go.P5Instance) {
	p.CreateCanvas(400, 400)
	p.NoStroke()
}

func draw(p *p5go.P5Instance) {
	// Gradient-like background
	for i := 0; i < 400; i++ {
		r := 220 - float64(i)*0.3
		g := 20 - float64(i)*0.05
		b := 60 - float64(i)*0.1
		if r < 0 { r = 0 }
		if g < 0 { g = 0 }
		if b < 0 { b = 0 }
		p.Fill(r, g, b)
		p.Rect(0, float64(i), 400, 1)
	}

	// Draw rotating Ruby gem (as a multi-faceted shape)
	p.Push()
	p.Translate(200, 200)
	p.Rotate(angle)
	p.Fill(255, 0, 0)
	p.BeginShape()
	for i := 0; i < 6; i++ {
		x := math.Cos(angle+float64(i)*math.Pi/3) * 100
		y := math.Sin(angle+float64(i)*math.Pi/3) * 100
		p.Vertex(x, y)
	}
	p.EndShape(p5go.Close)
	p.Pop()

	// Draw elegant floating code snippets with transparency
	p.Fill(255, 255, 255, 180)
	p.TextSize(18)
	p.TextAlign(p5go.Center, p5go.Center)
	p.Text("def elegant_ruby", 200, 50)
	p.Text("puts 'Shine bright like a gem'", 200, 350)
	p.Text("end", 200, 370)

	// Draw softly glowing, rotating dots around the gem
	for i := 0; i < 12; i++ {
		x := 200 + math.Cos(angle+float64(i)*math.Pi/6)*150
		y := 200 + math.Sin(angle+float64(i)*math.Pi/6)*150
		p.Fill(255, 255, 255, 150)
		p.Ellipse(x, y, 15, 15)
	}

	// Update angle for rotation
	angle += 0.01
}