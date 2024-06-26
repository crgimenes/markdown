package markdown

import (
	"fmt"
	"io"
	"strings"
)

type Element struct {
	Tag        string
	Raw        string
	SubElement []Element
}

func Parse(r io.Reader) (elements []Element, err error) {
	a, err := io.ReadAll(r)
	if err != nil {
		return
	}
	l := strings.Split(string(a), "\n\n")
	for k, v := range l {
		fmt.Println(k, v)
	}
	return
}
