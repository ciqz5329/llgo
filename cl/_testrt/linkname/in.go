package main

import (
	_ "unsafe"

	_ "github.com/goplus/llgo/cl/internal/linktarget"
	"github.com/goplus/llgo/internal/runtime/c"
)

//go:linkname print github.com/goplus/llgo/cl/internal/linktarget.F
func print(a, b, c, d *c.Char)

type m struct {
	s string
}

//go:linkname setInfo github.com/goplus/llgo/cl/internal/linktarget.(*m).setInfo
func setInfo(*m, string)

//go:linkname info github.com/goplus/llgo/cl/internal/linktarget.m.info
func info(m) string

func main() {
	print(c.Str("a"), c.Str("b"), c.Str("c"), c.Str("d"))
	print(c.Str("1"), c.Str("2"), c.Str("3"), c.Str("4"))
	var m m
	setInfo(&m, "hello")
	println(info(m))
}
