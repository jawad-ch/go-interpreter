package vm

import "github.com/jawad-ch/go-interpreter/object"

type Frame struct {
	fn          *object.CompiledFunction
	ip          int
	basePointer int
}

func NewFrame(fn *object.CompiledFunction, basPointer int) *Frame {
	f := &Frame{
		fn:          fn,
		ip:          -1,
		basePointer: basPointer,
	}

	return f
}
