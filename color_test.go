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
	url := "https://tra.bytedance.net/app/13/cases?keyWord=820&searchName=&tags=%7B%220%22%3A%5B%22%E5%B0%8F%E8%A7%86%E9%A2%91%E4%B8%93%E9%A1%B9%22%5D%2C%222%22%3A%5B%22android%22%5D%2C%224%22%3A%5B%5D%7D&caseStatus=0&order=0&priority=&current=1&follow=1&dateRange="
	render := Render(url)
	println(render)
}
