package vm

import (
	"github.com/jawad-ch/go-interpreter/code"
	"github.com/jawad-ch/go-interpreter/object"
)

type Frame struct {
	cl          *object.Closure
	ip          int
	basePointer int
}

func NewFrame(cl *object.Closure, basPointer int) *Frame {
	f := &Frame{
		cl:          cl,
		ip:          -1,
		basePointer: basPointer,
	}

	return f
}

func (f *Frame) Instructions() code.Instructions {
	return f.cl.Fn.Instructions
}
