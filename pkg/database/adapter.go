package database

import (
	"strings"
)

// Writer інтерфейс, який описує методи для запису даних
type Writer interface {
	Data() string
}

// MarkdownAdapter адаптер для перетворення HTML в Markdown
type MarkdownAdapter struct {
	Writer
}

// WriteData перетворює HTML в Markdown і записує дані
func (ma MarkdownAdapter) Data() string {
	markdownData := strings.ReplaceAll(ma.Writer.Data(), "<p>", "")
	markdownData = strings.ReplaceAll(markdownData, "</p>", "\n\n")
	markdownData = strings.ReplaceAll(markdownData, "<br>", "  \n")
	markdownData = strings.ReplaceAll(markdownData, "<strong>", "**")
	markdownData = strings.ReplaceAll(markdownData, "</strong>", "**")
	markdownData = strings.ReplaceAll(markdownData, "<i>", "*")
	markdownData = strings.ReplaceAll(markdownData, "</i>", "*")

	return markdownData
}
