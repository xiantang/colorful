package colorful

import (
	"testing"
)

func TestPrintList(t *testing.T) {
	type name struct {
		aaa int
	}
	names := make([]name, 0)
	for i := 0; i < 10; i++ {
		names = append(names, name{aaa: i})
	}
	render := Render(names)
	println(render)
}

func TestRenderURL(t *testing.T) {
	url := ""
	render := Render(url)
	println(render)
}
