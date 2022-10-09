package utils

import (
	"bytes"

	"github.com/yuin/goldmark"
)

func MarkdownToHTML(md string) string {
	var buf bytes.Buffer
	err := goldmark.Convert([]byte(md), &buf)
	if err != nil {
		panic(err)
	}
	return buf.String()
}
