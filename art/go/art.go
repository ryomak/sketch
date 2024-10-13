package art

import (
	"syscall/js"
)

type Interface interface {
	Title() any

	Description() any

	Generate(this js.Value, args []js.Value) any
}

func Do(name string, i Interface) {
	c := make(chan struct{}, 0)
	js.Global().Set(name, js.ValueOf(map[string]any{
		"title":       js.ValueOf(i.Title()),
		"description": js.ValueOf(i.Description()),
		"generate":    js.FuncOf(i.Generate),
	}))
	<-c
}