package ctemplate

import (
	"html/template"
)

func unescaped (s string) interface{} {
	return template.HTML(s)
}