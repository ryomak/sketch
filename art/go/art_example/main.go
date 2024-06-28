package main

import (
	"fmt"
	"math/rand"
	"syscall/js"

	"github.com/ryomak/sketch/art"
)

func main() {
	art.Do(Art{})
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
	if len(args) < 1 {
		fmt.Println("Error: Canvas ID not provided")
		return nil
	}

	canvasID := args[0].String()
	canvas := js.Global().Get("document").Call("getElementById", canvasID)
	if canvas.IsNull() {
		fmt.Printf("Error: Canvas with ID '%s' not found\n", canvasID)
		return nil
	}

	ctx := canvas.Call("getContext", "2d")

	width := canvas.Get("width").Int()
	height := canvas.Get("height").Int()

	// キャンバスをクリア
	ctx.Call("clearRect", 0, 0, width, height)

	// ここでアートを生成
	// 例：ランダムな色の円を描画
	for i := 0; i < 50; i++ {
		x := rand.Float64() * float64(width)
		y := rand.Float64() * float64(height)
		radius := rand.Float64() * 20

		ctx.Set("fillStyle", randomColor())
		ctx.Call("beginPath")
		ctx.Call("arc", x, y, radius, 0, 2*3.14159)
		ctx.Call("fill")
	}

	return nil
}

func randomColor() string {
	r := rand.Intn(256)
	g := rand.Intn(256)
	b := rand.Intn(256)
	return fmt.Sprintf("rgb(%d, %d, %d)", r, g, b)
}
